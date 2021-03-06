// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/cri (interfaces: APIs)

// Package mock_cri is a generated GoMock package.
package mock_cri

import (
	reflect "reflect"

	cri "github.com/aws/amazon-vpc-cni-k8s/pkg/cri"
	gomock "github.com/golang/mock/gomock"
)

// MockAPIs is a mock of APIs interface
type MockAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockAPIsMockRecorder
}

// MockAPIsMockRecorder is the mock recorder for MockAPIs
type MockAPIsMockRecorder struct {
	mock *MockAPIs
}

// NewMockAPIs creates a new mock instance
func NewMockAPIs(ctrl *gomock.Controller) *MockAPIs {
	mock := &MockAPIs{ctrl: ctrl}
	mock.recorder = &MockAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIs) EXPECT() *MockAPIsMockRecorder {
	return m.recorder
}

// GetRunningPodSandboxes mocks base method
func (m *MockAPIs) GetRunningPodSandboxes() (map[string]*cri.SandboxInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunningPodSandboxes")
	ret0, _ := ret[0].(map[string]*cri.SandboxInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunningPodSandboxes indicates an expected call of GetRunningPodSandboxes
func (mr *MockAPIsMockRecorder) GetRunningPodSandboxes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunningPodSandboxes", reflect.TypeOf((*MockAPIs)(nil).GetRunningPodSandboxes))
}
