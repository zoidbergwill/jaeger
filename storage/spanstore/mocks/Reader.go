// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import (
	"context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/jaegertracing/jaeger/model"
	spanstore "github.com/jaegertracing/jaeger/storage/spanstore"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// FindTraces provides a mock function with given fields: query
func (_m *Reader) FindTraces(ctx context.Context, query *spanstore.TraceQueryParameters) ([]*model.Trace, error) {
	ret := _m.Called(query)

	var r0 []*model.Trace
	if rf, ok := ret.Get(0).(func(*spanstore.TraceQueryParameters) []*model.Trace); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Trace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*spanstore.TraceQueryParameters) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperations provides a mock function with given fields: service
func (_m *Reader) GetOperations(ctx context.Context, service string) ([]string, error) {
	ret := _m.Called(service)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(service)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServices provides a mock function with given fields:
func (_m *Reader) GetServices(ctx context.Context) ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrace provides a mock function with given fields: traceID
func (_m *Reader) GetTrace(ctx context.Context, traceID model.TraceID) (*model.Trace, error) {
	ret := _m.Called(traceID)

	var r0 *model.Trace
	if rf, ok := ret.Get(0).(func(model.TraceID) *model.Trace); ok {
		r0 = rf(traceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Trace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.TraceID) error); ok {
		r1 = rf(traceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ spanstore.Reader = (*Reader)(nil)
