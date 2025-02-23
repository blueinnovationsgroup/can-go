// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/blueinnovationsgroup/can-go/pkg/canrunner (interfaces: Node,TransmittedMessage,ReceivedMessage,FrameTransmitter,FrameReceiver)

// Package mockcanrunner is a generated GoMock package.
package mockcanrunner

import (
	context "context"
	net "net"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	can "github.com/blueinnovationsgroup/can-go"
	canrunner "github.com/blueinnovationsgroup/can-go/pkg/canrunner"
	descriptor "github.com/blueinnovationsgroup/can-go/pkg/descriptor"
)

// MockNode is a mock of Node interface.
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode.
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance.
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockNode) Connect() (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockNodeMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNode)(nil).Connect))
}

// Descriptor mocks base method.
func (m *MockNode) Descriptor() *descriptor.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(*descriptor.Node)
	return ret0
}

// Descriptor indicates an expected call of Descriptor.
func (mr *MockNodeMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MockNode)(nil).Descriptor))
}

// Lock mocks base method.
func (m *MockNode) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock.
func (mr *MockNodeMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockNode)(nil).Lock))
}

// ReceivedMessage mocks base method.
func (m *MockNode) ReceivedMessage(arg0 uint32) (canrunner.ReceivedMessage, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivedMessage", arg0)
	ret0, _ := ret[0].(canrunner.ReceivedMessage)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ReceivedMessage indicates an expected call of ReceivedMessage.
func (mr *MockNodeMockRecorder) ReceivedMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedMessage", reflect.TypeOf((*MockNode)(nil).ReceivedMessage), arg0)
}

// TransmittedMessages mocks base method.
func (m *MockNode) TransmittedMessages() []canrunner.TransmittedMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransmittedMessages")
	ret0, _ := ret[0].([]canrunner.TransmittedMessage)
	return ret0
}

// TransmittedMessages indicates an expected call of TransmittedMessages.
func (mr *MockNodeMockRecorder) TransmittedMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransmittedMessages", reflect.TypeOf((*MockNode)(nil).TransmittedMessages))
}

// Unlock mocks base method.
func (m *MockNode) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock.
func (mr *MockNodeMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockNode)(nil).Unlock))
}

// MockTransmittedMessage is a mock of TransmittedMessage interface.
type MockTransmittedMessage struct {
	ctrl     *gomock.Controller
	recorder *MockTransmittedMessageMockRecorder
}

// MockTransmittedMessageMockRecorder is the mock recorder for MockTransmittedMessage.
type MockTransmittedMessageMockRecorder struct {
	mock *MockTransmittedMessage
}

// NewMockTransmittedMessage creates a new mock instance.
func NewMockTransmittedMessage(ctrl *gomock.Controller) *MockTransmittedMessage {
	mock := &MockTransmittedMessage{ctrl: ctrl}
	mock.recorder = &MockTransmittedMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransmittedMessage) EXPECT() *MockTransmittedMessageMockRecorder {
	return m.recorder
}

// BeforeTransmitHook mocks base method.
func (m *MockTransmittedMessage) BeforeTransmitHook() func(context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeTransmitHook")
	ret0, _ := ret[0].(func(context.Context) error)
	return ret0
}

// BeforeTransmitHook indicates an expected call of BeforeTransmitHook.
func (mr *MockTransmittedMessageMockRecorder) BeforeTransmitHook() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeTransmitHook", reflect.TypeOf((*MockTransmittedMessage)(nil).BeforeTransmitHook))
}

// Descriptor mocks base method.
func (m *MockTransmittedMessage) Descriptor() *descriptor.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(*descriptor.Message)
	return ret0
}

// Descriptor indicates an expected call of Descriptor.
func (mr *MockTransmittedMessageMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MockTransmittedMessage)(nil).Descriptor))
}

// Frame mocks base method.
func (m *MockTransmittedMessage) Frame() can.Frame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Frame")
	ret0, _ := ret[0].(can.Frame)
	return ret0
}

// Frame indicates an expected call of Frame.
func (mr *MockTransmittedMessageMockRecorder) Frame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Frame", reflect.TypeOf((*MockTransmittedMessage)(nil).Frame))
}

// IsCyclicTransmissionEnabled mocks base method.
func (m *MockTransmittedMessage) IsCyclicTransmissionEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCyclicTransmissionEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCyclicTransmissionEnabled indicates an expected call of IsCyclicTransmissionEnabled.
func (mr *MockTransmittedMessageMockRecorder) IsCyclicTransmissionEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCyclicTransmissionEnabled", reflect.TypeOf((*MockTransmittedMessage)(nil).IsCyclicTransmissionEnabled))
}

// MarshalFrame mocks base method.
func (m *MockTransmittedMessage) MarshalFrame() (can.Frame, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalFrame")
	ret0, _ := ret[0].(can.Frame)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalFrame indicates an expected call of MarshalFrame.
func (mr *MockTransmittedMessageMockRecorder) MarshalFrame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalFrame", reflect.TypeOf((*MockTransmittedMessage)(nil).MarshalFrame))
}

// Reset mocks base method.
func (m *MockTransmittedMessage) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockTransmittedMessageMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockTransmittedMessage)(nil).Reset))
}

// SetTransmitTime mocks base method.
func (m *MockTransmittedMessage) SetTransmitTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransmitTime", arg0)
}

// SetTransmitTime indicates an expected call of SetTransmitTime.
func (mr *MockTransmittedMessageMockRecorder) SetTransmitTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransmitTime", reflect.TypeOf((*MockTransmittedMessage)(nil).SetTransmitTime), arg0)
}

// String mocks base method.
func (m *MockTransmittedMessage) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockTransmittedMessageMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockTransmittedMessage)(nil).String))
}

// TransmitEventChan mocks base method.
func (m *MockTransmittedMessage) TransmitEventChan() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransmitEventChan")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// TransmitEventChan indicates an expected call of TransmitEventChan.
func (mr *MockTransmittedMessageMockRecorder) TransmitEventChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransmitEventChan", reflect.TypeOf((*MockTransmittedMessage)(nil).TransmitEventChan))
}

// UnmarshalFrame mocks base method.
func (m *MockTransmittedMessage) UnmarshalFrame(arg0 can.Frame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalFrame indicates an expected call of UnmarshalFrame.
func (mr *MockTransmittedMessageMockRecorder) UnmarshalFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalFrame", reflect.TypeOf((*MockTransmittedMessage)(nil).UnmarshalFrame), arg0)
}

// WakeUpChan mocks base method.
func (m *MockTransmittedMessage) WakeUpChan() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WakeUpChan")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// WakeUpChan indicates an expected call of WakeUpChan.
func (mr *MockTransmittedMessageMockRecorder) WakeUpChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WakeUpChan", reflect.TypeOf((*MockTransmittedMessage)(nil).WakeUpChan))
}

// MockReceivedMessage is a mock of ReceivedMessage interface.
type MockReceivedMessage struct {
	ctrl     *gomock.Controller
	recorder *MockReceivedMessageMockRecorder
}

// MockReceivedMessageMockRecorder is the mock recorder for MockReceivedMessage.
type MockReceivedMessageMockRecorder struct {
	mock *MockReceivedMessage
}

// NewMockReceivedMessage creates a new mock instance.
func NewMockReceivedMessage(ctrl *gomock.Controller) *MockReceivedMessage {
	mock := &MockReceivedMessage{ctrl: ctrl}
	mock.recorder = &MockReceivedMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceivedMessage) EXPECT() *MockReceivedMessageMockRecorder {
	return m.recorder
}

// AfterReceiveHook mocks base method.
func (m *MockReceivedMessage) AfterReceiveHook() func(context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterReceiveHook")
	ret0, _ := ret[0].(func(context.Context) error)
	return ret0
}

// AfterReceiveHook indicates an expected call of AfterReceiveHook.
func (mr *MockReceivedMessageMockRecorder) AfterReceiveHook() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterReceiveHook", reflect.TypeOf((*MockReceivedMessage)(nil).AfterReceiveHook))
}

// Descriptor mocks base method.
func (m *MockReceivedMessage) Descriptor() *descriptor.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(*descriptor.Message)
	return ret0
}

// Descriptor indicates an expected call of Descriptor.
func (mr *MockReceivedMessageMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MockReceivedMessage)(nil).Descriptor))
}

// Frame mocks base method.
func (m *MockReceivedMessage) Frame() can.Frame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Frame")
	ret0, _ := ret[0].(can.Frame)
	return ret0
}

// Frame indicates an expected call of Frame.
func (mr *MockReceivedMessageMockRecorder) Frame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Frame", reflect.TypeOf((*MockReceivedMessage)(nil).Frame))
}

// MarshalFrame mocks base method.
func (m *MockReceivedMessage) MarshalFrame() (can.Frame, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalFrame")
	ret0, _ := ret[0].(can.Frame)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalFrame indicates an expected call of MarshalFrame.
func (mr *MockReceivedMessageMockRecorder) MarshalFrame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalFrame", reflect.TypeOf((*MockReceivedMessage)(nil).MarshalFrame))
}

// Reset mocks base method.
func (m *MockReceivedMessage) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockReceivedMessageMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockReceivedMessage)(nil).Reset))
}

// SetReceiveTime mocks base method.
func (m *MockReceivedMessage) SetReceiveTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReceiveTime", arg0)
}

// SetReceiveTime indicates an expected call of SetReceiveTime.
func (mr *MockReceivedMessageMockRecorder) SetReceiveTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReceiveTime", reflect.TypeOf((*MockReceivedMessage)(nil).SetReceiveTime), arg0)
}

// String mocks base method.
func (m *MockReceivedMessage) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockReceivedMessageMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockReceivedMessage)(nil).String))
}

// UnmarshalFrame mocks base method.
func (m *MockReceivedMessage) UnmarshalFrame(arg0 can.Frame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalFrame indicates an expected call of UnmarshalFrame.
func (mr *MockReceivedMessageMockRecorder) UnmarshalFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalFrame", reflect.TypeOf((*MockReceivedMessage)(nil).UnmarshalFrame), arg0)
}

// MockFrameTransmitter is a mock of FrameTransmitter interface.
type MockFrameTransmitter struct {
	ctrl     *gomock.Controller
	recorder *MockFrameTransmitterMockRecorder
}

// MockFrameTransmitterMockRecorder is the mock recorder for MockFrameTransmitter.
type MockFrameTransmitterMockRecorder struct {
	mock *MockFrameTransmitter
}

// NewMockFrameTransmitter creates a new mock instance.
func NewMockFrameTransmitter(ctrl *gomock.Controller) *MockFrameTransmitter {
	mock := &MockFrameTransmitter{ctrl: ctrl}
	mock.recorder = &MockFrameTransmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrameTransmitter) EXPECT() *MockFrameTransmitterMockRecorder {
	return m.recorder
}

// TransmitFrame mocks base method.
func (m *MockFrameTransmitter) TransmitFrame(arg0 context.Context, arg1 can.Frame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransmitFrame", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransmitFrame indicates an expected call of TransmitFrame.
func (mr *MockFrameTransmitterMockRecorder) TransmitFrame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransmitFrame", reflect.TypeOf((*MockFrameTransmitter)(nil).TransmitFrame), arg0, arg1)
}

// MockFrameReceiver is a mock of FrameReceiver interface.
type MockFrameReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockFrameReceiverMockRecorder
}

// MockFrameReceiverMockRecorder is the mock recorder for MockFrameReceiver.
type MockFrameReceiverMockRecorder struct {
	mock *MockFrameReceiver
}

// NewMockFrameReceiver creates a new mock instance.
func NewMockFrameReceiver(ctrl *gomock.Controller) *MockFrameReceiver {
	mock := &MockFrameReceiver{ctrl: ctrl}
	mock.recorder = &MockFrameReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrameReceiver) EXPECT() *MockFrameReceiverMockRecorder {
	return m.recorder
}

// Err mocks base method.
func (m *MockFrameReceiver) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockFrameReceiverMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockFrameReceiver)(nil).Err))
}

// Frame mocks base method.
func (m *MockFrameReceiver) Frame() can.Frame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Frame")
	ret0, _ := ret[0].(can.Frame)
	return ret0
}

// Frame indicates an expected call of Frame.
func (mr *MockFrameReceiverMockRecorder) Frame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Frame", reflect.TypeOf((*MockFrameReceiver)(nil).Frame))
}

// Receive mocks base method.
func (m *MockFrameReceiver) Receive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *MockFrameReceiverMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockFrameReceiver)(nil).Receive))
}
