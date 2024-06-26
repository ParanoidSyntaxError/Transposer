// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	graph "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/graph"

	mock "github.com/stretchr/testify/mock"
)

// Discoverer is an autogenerated mock type for the Discoverer type
type Discoverer struct {
	mock.Mock
}

// Discover provides a mock function with given fields: ctx
func (_m *Discoverer) Discover(ctx context.Context) (graph.Graph, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Discover")
	}

	var r0 graph.Graph
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (graph.Graph, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) graph.Graph); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graph.Graph)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDiscoverer creates a new instance of Discoverer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDiscoverer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Discoverer {
	mock := &Discoverer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
