package mocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// DTMFSender is an autogenerated mock type for the DTMFSender type
type DTMFSender struct {
	mock.Mock
}

// SendDTMF provides a mock function with given fields: dtmf, opts
func (_m *DTMFSender) SendDTMF(dtmf string, opts *ari.DTMFOptions) {
	_m.Called(dtmf, opts)
}