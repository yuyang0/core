// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	enginetypes "github.com/projecteru2/core/engine/types"
	coretypes "github.com/projecteru2/core/types"

	mock "github.com/stretchr/testify/mock"

	pluginstypes "github.com/projecteru2/core/resource/plugins/types"

	types "github.com/projecteru2/core/resource/types"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AddNode provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Manager) AddNode(_a0 context.Context, _a1 string, _a2 types.Resources, _a3 *enginetypes.Info) (types.Resources, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 types.Resources
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Resources, *enginetypes.Info) types.Resources); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Resources)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.Resources, *enginetypes.Info) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Alloc provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Manager) Alloc(_a0 context.Context, _a1 string, _a2 int, _a3 types.Resources) ([]types.Resources, []types.Resources, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []types.Resources
	if rf, ok := ret.Get(0).(func(context.Context, string, int, types.Resources) []types.Resources); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Resources)
		}
	}

	var r1 []types.Resources
	if rf, ok := ret.Get(1).(func(context.Context, string, int, types.Resources) []types.Resources); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]types.Resources)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int, types.Resources) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetMetricsDescription provides a mock function with given fields: _a0
func (_m *Manager) GetMetricsDescription(_a0 context.Context) ([]*pluginstypes.MetricsDescription, error) {
	ret := _m.Called(_a0)

	var r0 []*pluginstypes.MetricsDescription
	if rf, ok := ret.Get(0).(func(context.Context) []*pluginstypes.MetricsDescription); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pluginstypes.MetricsDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMostIdleNode provides a mock function with given fields: _a0, _a1
func (_m *Manager) GetMostIdleNode(_a0 context.Context, _a1 []string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, []string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeMetrics provides a mock function with given fields: _a0, _a1
func (_m *Manager) GetNodeMetrics(_a0 context.Context, _a1 *coretypes.Node) ([]*pluginstypes.Metrics, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*pluginstypes.Metrics
	if rf, ok := ret.Get(0).(func(context.Context, *coretypes.Node) []*pluginstypes.Metrics); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pluginstypes.Metrics)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *coretypes.Node) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeResourceInfo provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Manager) GetNodeResourceInfo(_a0 context.Context, _a1 string, _a2 []*coretypes.Workload, _a3 bool) (types.Resources, types.Resources, []string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 types.Resources
	if rf, ok := ret.Get(0).(func(context.Context, string, []*coretypes.Workload, bool) types.Resources); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Resources)
		}
	}

	var r1 types.Resources
	if rf, ok := ret.Get(1).(func(context.Context, string, []*coretypes.Workload, bool) types.Resources); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(types.Resources)
		}
	}

	var r2 []string
	if rf, ok := ret.Get(2).(func(context.Context, string, []*coretypes.Workload, bool) []string); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]string)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, string, []*coretypes.Workload, bool) error); ok {
		r3 = rf(_a0, _a1, _a2, _a3)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetNodesDeployCapacity provides a mock function with given fields: _a0, _a1, _a2
func (_m *Manager) GetNodesDeployCapacity(_a0 context.Context, _a1 []string, _a2 types.Resources) (map[string]*pluginstypes.NodeDeployCapacity, int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 map[string]*pluginstypes.NodeDeployCapacity
	if rf, ok := ret.Get(0).(func(context.Context, []string, types.Resources) map[string]*pluginstypes.NodeDeployCapacity); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*pluginstypes.NodeDeployCapacity)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, []string, types.Resources) int); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, types.Resources) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Realloc provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Manager) Realloc(_a0 context.Context, _a1 string, _a2 types.Resources, _a3 types.Resources) (types.Resources, types.Resources, types.Resources, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 types.Resources
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Resources, types.Resources) types.Resources); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Resources)
		}
	}

	var r1 types.Resources
	if rf, ok := ret.Get(1).(func(context.Context, string, types.Resources, types.Resources) types.Resources); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(types.Resources)
		}
	}

	var r2 types.Resources
	if rf, ok := ret.Get(2).(func(context.Context, string, types.Resources, types.Resources) types.Resources); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(types.Resources)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, string, types.Resources, types.Resources) error); ok {
		r3 = rf(_a0, _a1, _a2, _a3)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Remap provides a mock function with given fields: _a0, _a1, _a2
func (_m *Manager) Remap(_a0 context.Context, _a1 string, _a2 []*coretypes.Workload) (map[string]types.Resources, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 map[string]types.Resources
	if rf, ok := ret.Get(0).(func(context.Context, string, []*coretypes.Workload) map[string]types.Resources); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]types.Resources)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []*coretypes.Workload) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveNode provides a mock function with given fields: _a0, _a1
func (_m *Manager) RemoveNode(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackAlloc provides a mock function with given fields: _a0, _a1, _a2
func (_m *Manager) RollbackAlloc(_a0 context.Context, _a1 string, _a2 []types.Resources) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []types.Resources) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackRealloc provides a mock function with given fields: _a0, _a1, _a2
func (_m *Manager) RollbackRealloc(_a0 context.Context, _a1 string, _a2 types.Resources) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Resources) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNodeResourceCapacity provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *Manager) SetNodeResourceCapacity(_a0 context.Context, _a1 string, _a2 types.Resources, _a3 types.Resources, _a4 bool, _a5 bool) (types.Resources, types.Resources, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 types.Resources
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Resources, types.Resources, bool, bool) types.Resources); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Resources)
		}
	}

	var r1 types.Resources
	if rf, ok := ret.Get(1).(func(context.Context, string, types.Resources, types.Resources, bool, bool) types.Resources); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(types.Resources)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, types.Resources, types.Resources, bool, bool) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetNodeResourceUsage provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *Manager) SetNodeResourceUsage(_a0 context.Context, _a1 string, _a2 types.Resources, _a3 types.Resources, _a4 []types.Resources, _a5 bool, _a6 bool) (types.Resources, types.Resources, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 types.Resources
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Resources, types.Resources, []types.Resources, bool, bool) types.Resources); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Resources)
		}
	}

	var r1 types.Resources
	if rf, ok := ret.Get(1).(func(context.Context, string, types.Resources, types.Resources, []types.Resources, bool, bool) types.Resources); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(types.Resources)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, types.Resources, types.Resources, []types.Resources, bool, bool) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t mockConstructorTestingTNewManager) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
