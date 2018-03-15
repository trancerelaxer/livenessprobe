// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/container-storage-interface/spec/lib/go/csi/v0 (interfaces: IdentityServer,ControllerServer,NodeServer)

package driver

import (
	context "context"

	v0 "github.com/container-storage-interface/spec/lib/go/csi/v0"
	gomock "github.com/golang/mock/gomock"
)

// Mock of IdentityServer interface
type MockIdentityServer struct {
	ctrl     *gomock.Controller
	recorder *_MockIdentityServerRecorder
}

// Recorder for MockIdentityServer (not exported)
type _MockIdentityServerRecorder struct {
	mock *MockIdentityServer
}

func NewMockIdentityServer(ctrl *gomock.Controller) *MockIdentityServer {
	mock := &MockIdentityServer{ctrl: ctrl}
	mock.recorder = &_MockIdentityServerRecorder{mock}
	return mock
}

func (_m *MockIdentityServer) EXPECT() *_MockIdentityServerRecorder {
	return _m.recorder
}

func (_m *MockIdentityServer) GetPluginCapabilities(_param0 context.Context, _param1 *v0.GetPluginCapabilitiesRequest) (*v0.GetPluginCapabilitiesResponse, error) {
	ret := _m.ctrl.Call(_m, "GetPluginCapabilities", _param0, _param1)
	ret0, _ := ret[0].(*v0.GetPluginCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIdentityServerRecorder) GetPluginCapabilities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPluginCapabilities", arg0, arg1)
}

func (_m *MockIdentityServer) GetPluginInfo(_param0 context.Context, _param1 *v0.GetPluginInfoRequest) (*v0.GetPluginInfoResponse, error) {
	ret := _m.ctrl.Call(_m, "GetPluginInfo", _param0, _param1)
	ret0, _ := ret[0].(*v0.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIdentityServerRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPluginInfo", arg0, arg1)
}

func (_m *MockIdentityServer) Probe(_param0 context.Context, _param1 *v0.ProbeRequest) (*v0.ProbeResponse, error) {
	ret := _m.ctrl.Call(_m, "Probe", _param0, _param1)
	ret0, _ := ret[0].(*v0.ProbeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIdentityServerRecorder) Probe(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Probe", arg0, arg1)
}

// Mock of ControllerServer interface
type MockControllerServer struct {
	ctrl     *gomock.Controller
	recorder *_MockControllerServerRecorder
}

// Recorder for MockControllerServer (not exported)
type _MockControllerServerRecorder struct {
	mock *MockControllerServer
}

func NewMockControllerServer(ctrl *gomock.Controller) *MockControllerServer {
	mock := &MockControllerServer{ctrl: ctrl}
	mock.recorder = &_MockControllerServerRecorder{mock}
	return mock
}

func (_m *MockControllerServer) EXPECT() *_MockControllerServerRecorder {
	return _m.recorder
}

func (_m *MockControllerServer) ControllerGetCapabilities(_param0 context.Context, _param1 *v0.ControllerGetCapabilitiesRequest) (*v0.ControllerGetCapabilitiesResponse, error) {
	ret := _m.ctrl.Call(_m, "ControllerGetCapabilities", _param0, _param1)
	ret0, _ := ret[0].(*v0.ControllerGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ControllerGetCapabilities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerGetCapabilities", arg0, arg1)
}

func (_m *MockControllerServer) ControllerPublishVolume(_param0 context.Context, _param1 *v0.ControllerPublishVolumeRequest) (*v0.ControllerPublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "ControllerPublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.ControllerPublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ControllerPublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerPublishVolume", arg0, arg1)
}

func (_m *MockControllerServer) ControllerUnpublishVolume(_param0 context.Context, _param1 *v0.ControllerUnpublishVolumeRequest) (*v0.ControllerUnpublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "ControllerUnpublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.ControllerUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ControllerUnpublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerUnpublishVolume", arg0, arg1)
}

func (_m *MockControllerServer) CreateVolume(_param0 context.Context, _param1 *v0.CreateVolumeRequest) (*v0.CreateVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.CreateVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) CreateVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateVolume", arg0, arg1)
}

func (_m *MockControllerServer) DeleteVolume(_param0 context.Context, _param1 *v0.DeleteVolumeRequest) (*v0.DeleteVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.DeleteVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) DeleteVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteVolume", arg0, arg1)
}

func (_m *MockControllerServer) GetCapacity(_param0 context.Context, _param1 *v0.GetCapacityRequest) (*v0.GetCapacityResponse, error) {
	ret := _m.ctrl.Call(_m, "GetCapacity", _param0, _param1)
	ret0, _ := ret[0].(*v0.GetCapacityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) GetCapacity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCapacity", arg0, arg1)
}

func (_m *MockControllerServer) ListVolumes(_param0 context.Context, _param1 *v0.ListVolumesRequest) (*v0.ListVolumesResponse, error) {
	ret := _m.ctrl.Call(_m, "ListVolumes", _param0, _param1)
	ret0, _ := ret[0].(*v0.ListVolumesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ListVolumes(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListVolumes", arg0, arg1)
}

func (_m *MockControllerServer) ValidateVolumeCapabilities(_param0 context.Context, _param1 *v0.ValidateVolumeCapabilitiesRequest) (*v0.ValidateVolumeCapabilitiesResponse, error) {
	ret := _m.ctrl.Call(_m, "ValidateVolumeCapabilities", _param0, _param1)
	ret0, _ := ret[0].(*v0.ValidateVolumeCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ValidateVolumeCapabilities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateVolumeCapabilities", arg0, arg1)
}

// Mock of NodeServer interface
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *_MockNodeServerRecorder
}

// Recorder for MockNodeServer (not exported)
type _MockNodeServerRecorder struct {
	mock *MockNodeServer
}

func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &_MockNodeServerRecorder{mock}
	return mock
}

func (_m *MockNodeServer) EXPECT() *_MockNodeServerRecorder {
	return _m.recorder
}

func (_m *MockNodeServer) NodeGetCapabilities(_param0 context.Context, _param1 *v0.NodeGetCapabilitiesRequest) (*v0.NodeGetCapabilitiesResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeGetCapabilities", _param0, _param1)
	ret0, _ := ret[0].(*v0.NodeGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeGetCapabilities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeGetCapabilities", arg0, arg1)
}

func (_m *MockNodeServer) NodeGetId(_param0 context.Context, _param1 *v0.NodeGetIdRequest) (*v0.NodeGetIdResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeGetId", _param0, _param1)
	ret0, _ := ret[0].(*v0.NodeGetIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeGetId(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeGetId", arg0, arg1)
}

func (_m *MockNodeServer) NodePublishVolume(_param0 context.Context, _param1 *v0.NodePublishVolumeRequest) (*v0.NodePublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "NodePublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.NodePublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodePublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodePublishVolume", arg0, arg1)
}

func (_m *MockNodeServer) NodeStageVolume(_param0 context.Context, _param1 *v0.NodeStageVolumeRequest) (*v0.NodeStageVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeStageVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.NodeStageVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeStageVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeStageVolume", arg0, arg1)
}

func (_m *MockNodeServer) NodeUnpublishVolume(_param0 context.Context, _param1 *v0.NodeUnpublishVolumeRequest) (*v0.NodeUnpublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeUnpublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.NodeUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeUnpublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeUnpublishVolume", arg0, arg1)
}

func (_m *MockNodeServer) NodeUnstageVolume(_param0 context.Context, _param1 *v0.NodeUnstageVolumeRequest) (*v0.NodeUnstageVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeUnstageVolume", _param0, _param1)
	ret0, _ := ret[0].(*v0.NodeUnstageVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeUnstageVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeUnstageVolume", arg0, arg1)
}