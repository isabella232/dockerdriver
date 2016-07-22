// This file was generated by counterfeiter
package voldriverfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/voldriver"
)

type FakeProvisioner struct {
	CreateStub        func(logger lager.Logger, createRequest voldriver.CreateRequest) voldriver.ErrorResponse
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		logger        lager.Logger
		createRequest voldriver.CreateRequest
	}
	createReturns struct {
		result1 voldriver.ErrorResponse
	}
	RemoveStub        func(logger lager.Logger, removeRequest voldriver.RemoveRequest) voldriver.ErrorResponse
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		logger        lager.Logger
		removeRequest voldriver.RemoveRequest
	}
	removeReturns struct {
		result1 voldriver.ErrorResponse
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProvisioner) Create(logger lager.Logger, createRequest voldriver.CreateRequest) voldriver.ErrorResponse {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		logger        lager.Logger
		createRequest voldriver.CreateRequest
	}{logger, createRequest})
	fake.recordInvocation("Create", []interface{}{logger, createRequest})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(logger, createRequest)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeProvisioner) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeProvisioner) CreateArgsForCall(i int) (lager.Logger, voldriver.CreateRequest) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].logger, fake.createArgsForCall[i].createRequest
}

func (fake *FakeProvisioner) CreateReturns(result1 voldriver.ErrorResponse) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 voldriver.ErrorResponse
	}{result1}
}

func (fake *FakeProvisioner) Remove(logger lager.Logger, removeRequest voldriver.RemoveRequest) voldriver.ErrorResponse {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		logger        lager.Logger
		removeRequest voldriver.RemoveRequest
	}{logger, removeRequest})
	fake.recordInvocation("Remove", []interface{}{logger, removeRequest})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(logger, removeRequest)
	} else {
		return fake.removeReturns.result1
	}
}

func (fake *FakeProvisioner) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeProvisioner) RemoveArgsForCall(i int) (lager.Logger, voldriver.RemoveRequest) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].logger, fake.removeArgsForCall[i].removeRequest
}

func (fake *FakeProvisioner) RemoveReturns(result1 voldriver.ErrorResponse) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 voldriver.ErrorResponse
	}{result1}
}

func (fake *FakeProvisioner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProvisioner) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ voldriver.Provisioner = new(FakeProvisioner)
