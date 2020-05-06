// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/18F/cf-cdn-service-broker/models"
)

type FakeRouteStore struct {
	CreateStub        func(*models.Route) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *models.Route
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	FindAllMatchingStub        func(models.Route) ([]models.Route, error)
	findAllMatchingMutex       sync.RWMutex
	findAllMatchingArgsForCall []struct {
		arg1 models.Route
	}
	findAllMatchingReturns struct {
		result1 []models.Route
		result2 error
	}
	findAllMatchingReturnsOnCall map[int]struct {
		result1 []models.Route
		result2 error
	}
	FindOneMatchingStub        func(models.Route) (models.Route, error)
	findOneMatchingMutex       sync.RWMutex
	findOneMatchingArgsForCall []struct {
		arg1 models.Route
	}
	findOneMatchingReturns struct {
		result1 models.Route
		result2 error
	}
	findOneMatchingReturnsOnCall map[int]struct {
		result1 models.Route
		result2 error
	}
	FindWithExpiringCertsStub        func() ([]models.Route, error)
	findWithExpiringCertsMutex       sync.RWMutex
	findWithExpiringCertsArgsForCall []struct {
	}
	findWithExpiringCertsReturns struct {
		result1 []models.Route
		result2 error
	}
	findWithExpiringCertsReturnsOnCall map[int]struct {
		result1 []models.Route
		result2 error
	}
	SaveStub        func(*models.Route) error
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		arg1 *models.Route
	}
	saveReturns struct {
		result1 error
	}
	saveReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouteStore) Create(arg1 *models.Route) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *models.Route
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeRouteStore) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRouteStore) CreateCalls(stub func(*models.Route) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeRouteStore) CreateArgsForCall(i int) *models.Route {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouteStore) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteStore) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteStore) FindAllMatching(arg1 models.Route) ([]models.Route, error) {
	fake.findAllMatchingMutex.Lock()
	ret, specificReturn := fake.findAllMatchingReturnsOnCall[len(fake.findAllMatchingArgsForCall)]
	fake.findAllMatchingArgsForCall = append(fake.findAllMatchingArgsForCall, struct {
		arg1 models.Route
	}{arg1})
	fake.recordInvocation("FindAllMatching", []interface{}{arg1})
	fake.findAllMatchingMutex.Unlock()
	if fake.FindAllMatchingStub != nil {
		return fake.FindAllMatchingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findAllMatchingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouteStore) FindAllMatchingCallCount() int {
	fake.findAllMatchingMutex.RLock()
	defer fake.findAllMatchingMutex.RUnlock()
	return len(fake.findAllMatchingArgsForCall)
}

func (fake *FakeRouteStore) FindAllMatchingCalls(stub func(models.Route) ([]models.Route, error)) {
	fake.findAllMatchingMutex.Lock()
	defer fake.findAllMatchingMutex.Unlock()
	fake.FindAllMatchingStub = stub
}

func (fake *FakeRouteStore) FindAllMatchingArgsForCall(i int) models.Route {
	fake.findAllMatchingMutex.RLock()
	defer fake.findAllMatchingMutex.RUnlock()
	argsForCall := fake.findAllMatchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouteStore) FindAllMatchingReturns(result1 []models.Route, result2 error) {
	fake.findAllMatchingMutex.Lock()
	defer fake.findAllMatchingMutex.Unlock()
	fake.FindAllMatchingStub = nil
	fake.findAllMatchingReturns = struct {
		result1 []models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteStore) FindAllMatchingReturnsOnCall(i int, result1 []models.Route, result2 error) {
	fake.findAllMatchingMutex.Lock()
	defer fake.findAllMatchingMutex.Unlock()
	fake.FindAllMatchingStub = nil
	if fake.findAllMatchingReturnsOnCall == nil {
		fake.findAllMatchingReturnsOnCall = make(map[int]struct {
			result1 []models.Route
			result2 error
		})
	}
	fake.findAllMatchingReturnsOnCall[i] = struct {
		result1 []models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteStore) FindOneMatching(arg1 models.Route) (models.Route, error) {
	fake.findOneMatchingMutex.Lock()
	ret, specificReturn := fake.findOneMatchingReturnsOnCall[len(fake.findOneMatchingArgsForCall)]
	fake.findOneMatchingArgsForCall = append(fake.findOneMatchingArgsForCall, struct {
		arg1 models.Route
	}{arg1})
	fake.recordInvocation("FindOneMatching", []interface{}{arg1})
	fake.findOneMatchingMutex.Unlock()
	if fake.FindOneMatchingStub != nil {
		return fake.FindOneMatchingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOneMatchingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouteStore) FindOneMatchingCallCount() int {
	fake.findOneMatchingMutex.RLock()
	defer fake.findOneMatchingMutex.RUnlock()
	return len(fake.findOneMatchingArgsForCall)
}

func (fake *FakeRouteStore) FindOneMatchingCalls(stub func(models.Route) (models.Route, error)) {
	fake.findOneMatchingMutex.Lock()
	defer fake.findOneMatchingMutex.Unlock()
	fake.FindOneMatchingStub = stub
}

func (fake *FakeRouteStore) FindOneMatchingArgsForCall(i int) models.Route {
	fake.findOneMatchingMutex.RLock()
	defer fake.findOneMatchingMutex.RUnlock()
	argsForCall := fake.findOneMatchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouteStore) FindOneMatchingReturns(result1 models.Route, result2 error) {
	fake.findOneMatchingMutex.Lock()
	defer fake.findOneMatchingMutex.Unlock()
	fake.FindOneMatchingStub = nil
	fake.findOneMatchingReturns = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteStore) FindOneMatchingReturnsOnCall(i int, result1 models.Route, result2 error) {
	fake.findOneMatchingMutex.Lock()
	defer fake.findOneMatchingMutex.Unlock()
	fake.FindOneMatchingStub = nil
	if fake.findOneMatchingReturnsOnCall == nil {
		fake.findOneMatchingReturnsOnCall = make(map[int]struct {
			result1 models.Route
			result2 error
		})
	}
	fake.findOneMatchingReturnsOnCall[i] = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteStore) FindWithExpiringCerts() ([]models.Route, error) {
	fake.findWithExpiringCertsMutex.Lock()
	ret, specificReturn := fake.findWithExpiringCertsReturnsOnCall[len(fake.findWithExpiringCertsArgsForCall)]
	fake.findWithExpiringCertsArgsForCall = append(fake.findWithExpiringCertsArgsForCall, struct {
	}{})
	fake.recordInvocation("FindWithExpiringCerts", []interface{}{})
	fake.findWithExpiringCertsMutex.Unlock()
	if fake.FindWithExpiringCertsStub != nil {
		return fake.FindWithExpiringCertsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findWithExpiringCertsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouteStore) FindWithExpiringCertsCallCount() int {
	fake.findWithExpiringCertsMutex.RLock()
	defer fake.findWithExpiringCertsMutex.RUnlock()
	return len(fake.findWithExpiringCertsArgsForCall)
}

func (fake *FakeRouteStore) FindWithExpiringCertsCalls(stub func() ([]models.Route, error)) {
	fake.findWithExpiringCertsMutex.Lock()
	defer fake.findWithExpiringCertsMutex.Unlock()
	fake.FindWithExpiringCertsStub = stub
}

func (fake *FakeRouteStore) FindWithExpiringCertsReturns(result1 []models.Route, result2 error) {
	fake.findWithExpiringCertsMutex.Lock()
	defer fake.findWithExpiringCertsMutex.Unlock()
	fake.FindWithExpiringCertsStub = nil
	fake.findWithExpiringCertsReturns = struct {
		result1 []models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteStore) FindWithExpiringCertsReturnsOnCall(i int, result1 []models.Route, result2 error) {
	fake.findWithExpiringCertsMutex.Lock()
	defer fake.findWithExpiringCertsMutex.Unlock()
	fake.FindWithExpiringCertsStub = nil
	if fake.findWithExpiringCertsReturnsOnCall == nil {
		fake.findWithExpiringCertsReturnsOnCall = make(map[int]struct {
			result1 []models.Route
			result2 error
		})
	}
	fake.findWithExpiringCertsReturnsOnCall[i] = struct {
		result1 []models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteStore) Save(arg1 *models.Route) error {
	fake.saveMutex.Lock()
	ret, specificReturn := fake.saveReturnsOnCall[len(fake.saveArgsForCall)]
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		arg1 *models.Route
	}{arg1})
	fake.recordInvocation("Save", []interface{}{arg1})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.saveReturns
	return fakeReturns.result1
}

func (fake *FakeRouteStore) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeRouteStore) SaveCalls(stub func(*models.Route) error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = stub
}

func (fake *FakeRouteStore) SaveArgsForCall(i int) *models.Route {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	argsForCall := fake.saveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouteStore) SaveReturns(result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteStore) SaveReturnsOnCall(i int, result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	if fake.saveReturnsOnCall == nil {
		fake.saveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.findAllMatchingMutex.RLock()
	defer fake.findAllMatchingMutex.RUnlock()
	fake.findOneMatchingMutex.RLock()
	defer fake.findOneMatchingMutex.RUnlock()
	fake.findWithExpiringCertsMutex.RLock()
	defer fake.findWithExpiringCertsMutex.RUnlock()
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRouteStore) recordInvocation(key string, args []interface{}) {
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

var _ models.RouteStoreInterface = new(FakeRouteStore)
