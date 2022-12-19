package models

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sanity-io/litter"

	"github.com/projecteru2/core/log"
	"github.com/projecteru2/core/resources/cpumem/types"
	"github.com/projecteru2/core/utils"
)

const NodeResourceInfoKey = "/resource/cpumem/%s"

// GetNodeResourceInfo .
func (c *CPUMem) GetNodeResourceInfo(ctx context.Context, node string, workloadResourceMap *types.WorkloadResourceArgsMap, fix bool) (*types.NodeResourceInfo, []string, error) {
	logger := log.WithFunc("resources.cpumem.GetNodeResourceInfo").WithField("node", node)
	resourceInfo, err := c.doGetNodeResourceInfo(ctx, node)
	if err != nil {
		logger.Error(ctx, err)
		return nil, nil, err
	}

	diffs := []string{}

	totalResourceArgs := &types.WorkloadResourceArgs{
		CPUMap:     types.CPUMap{},
		NUMAMemory: types.NUMAMemory{},
	}

	if workloadResourceMap != nil {
		for _, args := range *workloadResourceMap {
			totalResourceArgs.Add(args)
		}
	}

	totalResourceArgs.CPURequest = utils.Round(totalResourceArgs.CPURequest)
	totalCPUUsage := utils.Round(resourceInfo.Usage.CPU)
	if totalResourceArgs.CPURequest != totalCPUUsage {
		diffs = append(diffs, fmt.Sprintf("node.CPUUsed != sum(workload.CPURequest): %.2f != %.2f", totalCPUUsage, totalResourceArgs.CPURequest))
	}

	for cpu := range resourceInfo.Capacity.CPUMap {
		if totalResourceArgs.CPUMap[cpu] != resourceInfo.Usage.CPUMap[cpu] {
			diffs = append(diffs, fmt.Sprintf("node.CPUMap[%+v] != sum(workload.CPUMap[%+v]): %+v != %+v", cpu, cpu, resourceInfo.Usage.CPUMap[cpu], totalResourceArgs.CPUMap[cpu]))
		}
	}

	for numaNodeID := range resourceInfo.Capacity.NUMAMemory {
		if totalResourceArgs.NUMAMemory[numaNodeID] != resourceInfo.Usage.NUMAMemory[numaNodeID] {
			diffs = append(diffs, fmt.Sprintf("node.NUMAMemory[%+v] != sum(workload.NUMAMemory[%+v]: %+v != %+v)", numaNodeID, numaNodeID, resourceInfo.Usage.NUMAMemory[numaNodeID], totalResourceArgs.NUMAMemory[numaNodeID]))
		}
	}

	if resourceInfo.Usage.Memory != totalResourceArgs.MemoryRequest {
		diffs = append(diffs, fmt.Sprintf("node.MemoryUsed != sum(workload.MemoryRequest): %d != %d", resourceInfo.Usage.Memory, totalResourceArgs.MemoryRequest))
	}

	if fix {
		resourceInfo.Usage = &types.NodeResourceArgs{
			CPU:        totalResourceArgs.CPURequest,
			CPUMap:     totalResourceArgs.CPUMap,
			Memory:     totalResourceArgs.MemoryRequest,
			NUMAMemory: totalResourceArgs.NUMAMemory,
		}
		if err = c.doSetNodeResourceInfo(ctx, node, resourceInfo); err != nil {
			log.WithFunc("resources.cpumem.GetNodeResourceInfo").Error(ctx, err)
			diffs = append(diffs, err.Error())
		}
	}

	return resourceInfo, diffs, nil
}

// SetNodeResourceUsage .
func (c *CPUMem) SetNodeResourceUsage(ctx context.Context, node string, nodeResourceOpts *types.NodeResourceOpts, nodeResourceArgs *types.NodeResourceArgs, workloadResourceArgs []*types.WorkloadResourceArgs, delta bool, incr bool) (before *types.NodeResourceArgs, after *types.NodeResourceArgs, err error) {
	resourceInfo, err := c.doGetNodeResourceInfo(ctx, node)
	if err != nil {
		log.WithFunc("resources.cpumem.SetNodeResourceInfo").WithField("node", node).Error(ctx, err)
		return nil, nil, err
	}

	before = resourceInfo.Usage.DeepCopy()
	resourceInfo.Usage = c.calculateNodeResourceArgs(resourceInfo.Usage, nodeResourceOpts, nodeResourceArgs, workloadResourceArgs, delta, incr)

	if err := c.doSetNodeResourceInfo(ctx, node, resourceInfo); err != nil {
		return nil, nil, err
	}
	return before, resourceInfo.Usage, nil
}

// SetNodeResourceCapacity .
func (c *CPUMem) SetNodeResourceCapacity(ctx context.Context, node string, nodeResourceOpts *types.NodeResourceOpts, nodeResourceArgs *types.NodeResourceArgs, delta bool, incr bool) (before *types.NodeResourceArgs, after *types.NodeResourceArgs, err error) {
	resourceInfo, err := c.doGetNodeResourceInfo(ctx, node)
	if err != nil {
		log.WithFunc("resources.cpumem.SetNodeResourceCapacity").WithField("node", node).Error(ctx, err)
		return nil, nil, err
	}

	before = resourceInfo.Capacity.DeepCopy()
	if !delta && nodeResourceOpts != nil {
		nodeResourceOpts.SkipEmpty(resourceInfo.Capacity)
	}

	resourceInfo.Capacity = c.calculateNodeResourceArgs(resourceInfo.Capacity, nodeResourceOpts, nodeResourceArgs, nil, delta, incr)

	// add new cpu
	for cpu := range resourceInfo.Capacity.CPUMap {
		_, ok := resourceInfo.Usage.CPUMap[cpu]
		if !ok {
			resourceInfo.Usage.CPUMap[cpu] = 0
		}
	}

	// delete cpus with no pieces
	resourceInfo.RemoveEmptyCores()

	if err := c.doSetNodeResourceInfo(ctx, node, resourceInfo); err != nil {
		log.WithFunc("resources.cpumem.SetNodeResourceCapacity").WithField("node", node).Errorf(ctx, err, "resource info %+v", litter.Sdump(resourceInfo))
		return nil, nil, err
	}
	return before, resourceInfo.Capacity, nil
}

// SetNodeResourceInfo .
func (c *CPUMem) SetNodeResourceInfo(ctx context.Context, node string, resourceCapacity *types.NodeResourceArgs, resourceUsage *types.NodeResourceArgs) error {
	resourceInfo := &types.NodeResourceInfo{
		Capacity: resourceCapacity,
		Usage:    resourceUsage,
	}

	err := c.doSetNodeResourceInfo(ctx, node, resourceInfo)
	if err != nil {
		log.WithFunc("resources.cpumem.SetNodeResourceInfo").WithField("node", node).Errorf(ctx, err, "resource info %+v", litter.Sdump(resourceInfo))
	}
	return err
}

func (c *CPUMem) doGetNodeResourceInfo(ctx context.Context, node string) (*types.NodeResourceInfo, error) {
	resourceInfo := &types.NodeResourceInfo{}
	resp, err := c.store.GetOne(ctx, fmt.Sprintf(NodeResourceInfoKey, node))
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resp.Value, resourceInfo); err != nil {
		return nil, err
	}
	return resourceInfo, nil
}

func (c *CPUMem) doSetNodeResourceInfo(ctx context.Context, node string, resourceInfo *types.NodeResourceInfo) error {
	if err := resourceInfo.Validate(); err != nil {
		return err
	}

	data, err := json.Marshal(resourceInfo)
	if err != nil {
		return err
	}

	if _, err = c.store.Put(ctx, fmt.Sprintf(NodeResourceInfoKey, node), string(data)); err != nil {
		return err
	}
	return nil
}

// calculateNodeResourceArgs priority: node resource opts > node resource args > workload resource args list
func (c *CPUMem) calculateNodeResourceArgs(origin *types.NodeResourceArgs, nodeResourceOpts *types.NodeResourceOpts, nodeResourceArgs *types.NodeResourceArgs, workloadResourceArgs []*types.WorkloadResourceArgs, delta bool, incr bool) (res *types.NodeResourceArgs) {
	if origin == nil || !delta {
		res = (&types.NodeResourceArgs{}).DeepCopy()
	} else {
		res = origin.DeepCopy()
	}

	if nodeResourceOpts != nil {
		nodeResourceArgs := &types.NodeResourceArgs{
			CPU:        float64(len(nodeResourceOpts.CPUMap)),
			CPUMap:     nodeResourceOpts.CPUMap,
			Memory:     nodeResourceOpts.Memory,
			NUMAMemory: nodeResourceOpts.NUMAMemory,
			NUMA:       nodeResourceOpts.NUMA,
		}

		if incr {
			res.Add(nodeResourceArgs)
		} else {
			res.Sub(nodeResourceArgs)
		}
		return res
	}

	if nodeResourceArgs != nil {
		if incr {
			res.Add(nodeResourceArgs)
		} else {
			res.Sub(nodeResourceArgs)
		}
		return res
	}

	for _, args := range workloadResourceArgs {
		nodeResourceArgs := &types.NodeResourceArgs{
			CPU:        args.CPURequest,
			CPUMap:     args.CPUMap,
			NUMAMemory: args.NUMAMemory,
			Memory:     args.MemoryRequest,
		}
		if incr {
			res.Add(nodeResourceArgs)
		} else {
			res.Sub(nodeResourceArgs)
		}
	}
	return res
}
