package mocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// MailboxHandle is an autogenerated mock type for the MailboxHandle type
type MailboxHandle struct {
	mock.Mock
}

// Data provides a mock function with given fields:
func (_m *MailboxHandle) Data() (*ari.MailboxData, error) {
	ret := _m.Called()

	var r0 *ari.MailboxData
	if rf, ok := ret.Get(0).(func() *ari.MailboxData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.MailboxData)
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

// Delete provides a mock function with given fields:
func (_m *MailboxHandle) Delete() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *MailboxHandle) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Update provides a mock function with given fields: oldMessages, newMessages
func (_m *MailboxHandle) Update(oldMessages int, newMessages int) error {
	ret := _m.Called(oldMessages, newMessages)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(oldMessages, newMessages)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}