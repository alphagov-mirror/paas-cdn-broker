// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/18F/cf-cdn-service-broker/config"
	"github.com/18F/cf-cdn-service-broker/lego/acme"
	"github.com/18F/cf-cdn-service-broker/models"
	"github.com/18F/cf-cdn-service-broker/utils"
)

type FakeAcmeClientProvider struct {
	GetDNS01ClientStub        func(*utils.User, config.Settings) (acme.ClientInterface, error)
	getDNS01ClientMutex       sync.RWMutex
	getDNS01ClientArgsForCall []struct {
		arg1 *utils.User
		arg2 config.Settings
	}
	getDNS01ClientReturns struct {
		result1 acme.ClientInterface
		result2 error
	}
	getDNS01ClientReturnsOnCall map[int]struct {
		result1 acme.ClientInterface
		result2 error
	}
	GetHTTP01ClientStub        func(*utils.User, config.Settings) (acme.ClientInterface, error)
	getHTTP01ClientMutex       sync.RWMutex
	getHTTP01ClientArgsForCall []struct {
		arg1 *utils.User
		arg2 config.Settings
	}
	getHTTP01ClientReturns struct {
		result1 acme.ClientInterface
		result2 error
	}
	getHTTP01ClientReturnsOnCall map[int]struct {
		result1 acme.ClientInterface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAcmeClientProvider) GetDNS01Client(arg1 *utils.User, arg2 config.Settings) (acme.ClientInterface, error) {
	fake.getDNS01ClientMutex.Lock()
	ret, specificReturn := fake.getDNS01ClientReturnsOnCall[len(fake.getDNS01ClientArgsForCall)]
	fake.getDNS01ClientArgsForCall = append(fake.getDNS01ClientArgsForCall, struct {
		arg1 *utils.User
		arg2 config.Settings
	}{arg1, arg2})
	fake.recordInvocation("GetDNS01Client", []interface{}{arg1, arg2})
	fake.getDNS01ClientMutex.Unlock()
	if fake.GetDNS01ClientStub != nil {
		return fake.GetDNS01ClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getDNS01ClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAcmeClientProvider) GetDNS01ClientCallCount() int {
	fake.getDNS01ClientMutex.RLock()
	defer fake.getDNS01ClientMutex.RUnlock()
	return len(fake.getDNS01ClientArgsForCall)
}

func (fake *FakeAcmeClientProvider) GetDNS01ClientCalls(stub func(*utils.User, config.Settings) (acme.ClientInterface, error)) {
	fake.getDNS01ClientMutex.Lock()
	defer fake.getDNS01ClientMutex.Unlock()
	fake.GetDNS01ClientStub = stub
}

func (fake *FakeAcmeClientProvider) GetDNS01ClientArgsForCall(i int) (*utils.User, config.Settings) {
	fake.getDNS01ClientMutex.RLock()
	defer fake.getDNS01ClientMutex.RUnlock()
	argsForCall := fake.getDNS01ClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAcmeClientProvider) GetDNS01ClientReturns(result1 acme.ClientInterface, result2 error) {
	fake.getDNS01ClientMutex.Lock()
	defer fake.getDNS01ClientMutex.Unlock()
	fake.GetDNS01ClientStub = nil
	fake.getDNS01ClientReturns = struct {
		result1 acme.ClientInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeAcmeClientProvider) GetDNS01ClientReturnsOnCall(i int, result1 acme.ClientInterface, result2 error) {
	fake.getDNS01ClientMutex.Lock()
	defer fake.getDNS01ClientMutex.Unlock()
	fake.GetDNS01ClientStub = nil
	if fake.getDNS01ClientReturnsOnCall == nil {
		fake.getDNS01ClientReturnsOnCall = make(map[int]struct {
			result1 acme.ClientInterface
			result2 error
		})
	}
	fake.getDNS01ClientReturnsOnCall[i] = struct {
		result1 acme.ClientInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeAcmeClientProvider) GetHTTP01Client(arg1 *utils.User, arg2 config.Settings) (acme.ClientInterface, error) {
	fake.getHTTP01ClientMutex.Lock()
	ret, specificReturn := fake.getHTTP01ClientReturnsOnCall[len(fake.getHTTP01ClientArgsForCall)]
	fake.getHTTP01ClientArgsForCall = append(fake.getHTTP01ClientArgsForCall, struct {
		arg1 *utils.User
		arg2 config.Settings
	}{arg1, arg2})
	fake.recordInvocation("GetHTTP01Client", []interface{}{arg1, arg2})
	fake.getHTTP01ClientMutex.Unlock()
	if fake.GetHTTP01ClientStub != nil {
		return fake.GetHTTP01ClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getHTTP01ClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAcmeClientProvider) GetHTTP01ClientCallCount() int {
	fake.getHTTP01ClientMutex.RLock()
	defer fake.getHTTP01ClientMutex.RUnlock()
	return len(fake.getHTTP01ClientArgsForCall)
}

func (fake *FakeAcmeClientProvider) GetHTTP01ClientCalls(stub func(*utils.User, config.Settings) (acme.ClientInterface, error)) {
	fake.getHTTP01ClientMutex.Lock()
	defer fake.getHTTP01ClientMutex.Unlock()
	fake.GetHTTP01ClientStub = stub
}

func (fake *FakeAcmeClientProvider) GetHTTP01ClientArgsForCall(i int) (*utils.User, config.Settings) {
	fake.getHTTP01ClientMutex.RLock()
	defer fake.getHTTP01ClientMutex.RUnlock()
	argsForCall := fake.getHTTP01ClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAcmeClientProvider) GetHTTP01ClientReturns(result1 acme.ClientInterface, result2 error) {
	fake.getHTTP01ClientMutex.Lock()
	defer fake.getHTTP01ClientMutex.Unlock()
	fake.GetHTTP01ClientStub = nil
	fake.getHTTP01ClientReturns = struct {
		result1 acme.ClientInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeAcmeClientProvider) GetHTTP01ClientReturnsOnCall(i int, result1 acme.ClientInterface, result2 error) {
	fake.getHTTP01ClientMutex.Lock()
	defer fake.getHTTP01ClientMutex.Unlock()
	fake.GetHTTP01ClientStub = nil
	if fake.getHTTP01ClientReturnsOnCall == nil {
		fake.getHTTP01ClientReturnsOnCall = make(map[int]struct {
			result1 acme.ClientInterface
			result2 error
		})
	}
	fake.getHTTP01ClientReturnsOnCall[i] = struct {
		result1 acme.ClientInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeAcmeClientProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDNS01ClientMutex.RLock()
	defer fake.getDNS01ClientMutex.RUnlock()
	fake.getHTTP01ClientMutex.RLock()
	defer fake.getHTTP01ClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAcmeClientProvider) recordInvocation(key string, args []interface{}) {
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

var _ models.AcmeClientProviderInterface = new(FakeAcmeClientProvider)
