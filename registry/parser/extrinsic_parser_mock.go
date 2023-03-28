// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package parser

import (
	registry "github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	mock "github.com/stretchr/testify/mock"
)

// ExtrinsicParserMock is an autogenerated mock type for the ExtrinsicParser type
type ExtrinsicParserMock struct {
	mock.Mock
}

// ParseExtrinsics provides a mock function with given fields: callRegistry, block
func (_m *ExtrinsicParserMock) ParseExtrinsics(callRegistry registry.CallRegistry, block *types.SignedBlock) ([]*Extrinsic, error) {
	ret := _m.Called(callRegistry, block)

	var r0 []*Extrinsic
	if rf, ok := ret.Get(0).(func(registry.CallRegistry, *types.SignedBlock) []*Extrinsic); ok {
		r0 = rf(callRegistry, block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Extrinsic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(registry.CallRegistry, *types.SignedBlock) error); ok {
		r1 = rf(callRegistry, block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewExtrinsicParserMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewExtrinsicParserMock creates a new instance of ExtrinsicParserMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExtrinsicParserMock(t NewExtrinsicParserMockT) *ExtrinsicParserMock {
	mock := &ExtrinsicParserMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
