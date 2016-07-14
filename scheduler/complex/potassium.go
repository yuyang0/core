package complexscheduler

import (
	"fmt"

	"github.com/coreos/etcd/client"
	"gitlab.ricebook.net/platform/core/lock"
	"gitlab.ricebook.net/platform/core/types"
)

type potassium struct {
	*lock.Mutex
}

func NewPotassim(config types.Config) (*potassium, error) {
	if len(config.EtcdMachines) == 0 {
		return nil, fmt.Errorf("ETCD must be set")
	}

	cli, err := client.New(client.Config{Endpoints: config.EtcdMachines})
	if err != nil {
		return nil, err
	}

	mu := lock.NewMutex(client.NewKeysAPI(cli), config.Scheduler.EtcdLockKey, config.Scheduler.EtcdLockTTL)
	if mu == nil {
		return nil, fmt.Errorf("cannot init mutex")
	}
	return &potassium{mu}, nil
}

func (m *potassium) RandomNode(nodes map[string]types.CPUMap) (string, error) {
	m.Lock()
	defer m.Unlock()

	var nodename string
	if len(nodes) == 0 {
		return nodename, fmt.Errorf("No nodes provide to choose one")
	}
	max := 0
	for name, cpumap := range nodes {
		total := cpumap.Total()
		if total > max {
			max = total
			nodename = name
		}
	}

	// doesn't matter if max is still 0
	// which means no resource available
	return nodename, nil
}

func (m *potassium) SelectNodes(nodes map[string]types.CPUMap, quota float64, num int) (map[string][]types.CPUMap, error) {
	m.Lock()
	defer m.Unlock()

	result := make(map[string][]types.CPUMap)
	if len(nodes) == 0 {
		return result, fmt.Errorf("No nodes provide to choose some")
	}

	// all core could be shared
	// suppose each core has 10 coreShare
	// TODO: change it to be control by parameters
	result = AveragePlan(quota, nodes, num, -1, 10)
	if result == nil {
		return nil, fmt.Errorf("Not enough resource")
	}
	return result, nil
}
