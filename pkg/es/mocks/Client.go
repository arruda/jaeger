// Code generated by mockery v1.0.0

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package mocks

import es "github.com/uber/jaeger/pkg/es"
import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateIndex provides a mock function with given fields: index
func (_m *Client) CreateIndex(index string) es.IndicesCreateService {
	ret := _m.Called(index)

	var r0 es.IndicesCreateService
	if rf, ok := ret.Get(0).(func(string) es.IndicesCreateService); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndicesCreateService)
		}
	}

	return r0
}

// Index provides a mock function with given fields:
func (_m *Client) Index() es.IndexService {
	ret := _m.Called()

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func() es.IndexService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Search provides a mock function with given fields: indices
func (_m *Client) Search(indices ...string) es.SearchService {
	_va := make([]interface{}, len(indices))
	for _i := range indices {
		_va[_i] = indices[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 es.SearchService
	if rf, ok := ret.Get(0).(func(...string) es.SearchService); ok {
		r0 = rf(indices...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.SearchService)
		}
	}

	return r0
}