// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: Unpacker)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_unpacker_test.go github.com/quic-go/quic-go Unpacker
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	protocol "github.com/quic-go-no-crypto/quic-go-no-crypto/internal/protocol"
	wire "github.com/quic-go-no-crypto/quic-go-no-crypto/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockUnpacker is a mock of Unpacker interface.
type MockUnpacker struct {
	ctrl     *gomock.Controller
	recorder *MockUnpackerMockRecorder
}

// MockUnpackerMockRecorder is the mock recorder for MockUnpacker.
type MockUnpackerMockRecorder struct {
	mock *MockUnpacker
}

// NewMockUnpacker creates a new mock instance.
func NewMockUnpacker(ctrl *gomock.Controller) *MockUnpacker {
	mock := &MockUnpacker{ctrl: ctrl}
	mock.recorder = &MockUnpackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnpacker) EXPECT() *MockUnpackerMockRecorder {
	return m.recorder
}

// UnpackLongHeader mocks base method.
func (m *MockUnpacker) UnpackLongHeader(arg0 *wire.Header, arg1 time.Time, arg2 []byte, arg3 protocol.Version) (*unpackedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackLongHeader", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*unpackedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpackLongHeader indicates an expected call of UnpackLongHeader.
func (mr *MockUnpackerMockRecorder) UnpackLongHeader(arg0, arg1, arg2, arg3 any) *UnpackerUnpackLongHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackLongHeader", reflect.TypeOf((*MockUnpacker)(nil).UnpackLongHeader), arg0, arg1, arg2, arg3)
	return &UnpackerUnpackLongHeaderCall{Call: call}
}

// UnpackerUnpackLongHeaderCall wrap *gomock.Call
type UnpackerUnpackLongHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UnpackerUnpackLongHeaderCall) Return(arg0 *unpackedPacket, arg1 error) *UnpackerUnpackLongHeaderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UnpackerUnpackLongHeaderCall) Do(f func(*wire.Header, time.Time, []byte, protocol.Version) (*unpackedPacket, error)) *UnpackerUnpackLongHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UnpackerUnpackLongHeaderCall) DoAndReturn(f func(*wire.Header, time.Time, []byte, protocol.Version) (*unpackedPacket, error)) *UnpackerUnpackLongHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UnpackShortHeader mocks base method.
func (m *MockUnpacker) UnpackShortHeader(arg0 time.Time, arg1 []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackShortHeader", arg0, arg1)
	ret0, _ := ret[0].(protocol.PacketNumber)
	ret1, _ := ret[1].(protocol.PacketNumberLen)
	ret2, _ := ret[2].(protocol.KeyPhaseBit)
	ret3, _ := ret[3].([]byte)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// UnpackShortHeader indicates an expected call of UnpackShortHeader.
func (mr *MockUnpackerMockRecorder) UnpackShortHeader(arg0, arg1 any) *UnpackerUnpackShortHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackShortHeader", reflect.TypeOf((*MockUnpacker)(nil).UnpackShortHeader), arg0, arg1)
	return &UnpackerUnpackShortHeaderCall{Call: call}
}

// UnpackerUnpackShortHeaderCall wrap *gomock.Call
type UnpackerUnpackShortHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UnpackerUnpackShortHeaderCall) Return(arg0 protocol.PacketNumber, arg1 protocol.PacketNumberLen, arg2 protocol.KeyPhaseBit, arg3 []byte, arg4 error) *UnpackerUnpackShortHeaderCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3, arg4)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UnpackerUnpackShortHeaderCall) Do(f func(time.Time, []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error)) *UnpackerUnpackShortHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UnpackerUnpackShortHeaderCall) DoAndReturn(f func(time.Time, []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error)) *UnpackerUnpackShortHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
