// This file was generated by counterfeiter
package radarfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/radar"
)

type FakeRadarDB struct {
	GetPipelineNameStub        func() string
	getPipelineNameMutex       sync.RWMutex
	getPipelineNameArgsForCall []struct{}
	getPipelineNameReturns     struct {
		result1 string
	}
	GetPipelineIDStub        func() int
	getPipelineIDMutex       sync.RWMutex
	getPipelineIDArgsForCall []struct{}
	getPipelineIDReturns     struct {
		result1 int
	}
	ScopedNameStub        func(string) string
	scopedNameMutex       sync.RWMutex
	scopedNameArgsForCall []struct {
		arg1 string
	}
	scopedNameReturns struct {
		result1 string
	}
	TeamIDStub        func() int
	teamIDMutex       sync.RWMutex
	teamIDArgsForCall []struct{}
	teamIDReturns     struct {
		result1 int
	}
	ConfigStub        func() atc.Config
	configMutex       sync.RWMutex
	configArgsForCall []struct{}
	configReturns     struct {
		result1 atc.Config
	}
	IsPausedStub        func() (bool, error)
	isPausedMutex       sync.RWMutex
	isPausedArgsForCall []struct{}
	isPausedReturns     struct {
		result1 bool
		result2 error
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct{}
	reloadReturns     struct {
		result1 bool
		result2 error
	}
	GetLatestVersionedResourceStub        func(resourceName string) (db.SavedVersionedResource, bool, error)
	getLatestVersionedResourceMutex       sync.RWMutex
	getLatestVersionedResourceArgsForCall []struct {
		resourceName string
	}
	getLatestVersionedResourceReturns struct {
		result1 db.SavedVersionedResource
		result2 bool
		result3 error
	}
	GetResourceStub        func(resourceName string) (db.SavedResource, bool, error)
	getResourceMutex       sync.RWMutex
	getResourceArgsForCall []struct {
		resourceName string
	}
	getResourceReturns struct {
		result1 db.SavedResource
		result2 bool
		result3 error
	}
	GetResourceTypeStub        func(resourceTypeName string) (db.SavedResourceType, bool, error)
	getResourceTypeMutex       sync.RWMutex
	getResourceTypeArgsForCall []struct {
		resourceTypeName string
	}
	getResourceTypeReturns struct {
		result1 db.SavedResourceType
		result2 bool
		result3 error
	}
	PauseResourceStub        func(resourceName string) error
	pauseResourceMutex       sync.RWMutex
	pauseResourceArgsForCall []struct {
		resourceName string
	}
	pauseResourceReturns struct {
		result1 error
	}
	UnpauseResourceStub        func(resourceName string) error
	unpauseResourceMutex       sync.RWMutex
	unpauseResourceArgsForCall []struct {
		resourceName string
	}
	unpauseResourceReturns struct {
		result1 error
	}
	SaveResourceVersionsStub        func(atc.ResourceConfig, []atc.Version) error
	saveResourceVersionsMutex       sync.RWMutex
	saveResourceVersionsArgsForCall []struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}
	saveResourceVersionsReturns struct {
		result1 error
	}
	SaveResourceTypeVersionStub        func(atc.ResourceType, atc.Version) error
	saveResourceTypeVersionMutex       sync.RWMutex
	saveResourceTypeVersionArgsForCall []struct {
		arg1 atc.ResourceType
		arg2 atc.Version
	}
	saveResourceTypeVersionReturns struct {
		result1 error
	}
	SetResourceCheckErrorStub        func(resource db.SavedResource, err error) error
	setResourceCheckErrorMutex       sync.RWMutex
	setResourceCheckErrorArgsForCall []struct {
		resource db.SavedResource
		err      error
	}
	setResourceCheckErrorReturns struct {
		result1 error
	}
	AcquireResourceCheckingLockStub        func(logger lager.Logger, resource db.SavedResource, interval time.Duration, immediate bool) (db.Lock, bool, error)
	acquireResourceCheckingLockMutex       sync.RWMutex
	acquireResourceCheckingLockArgsForCall []struct {
		logger    lager.Logger
		resource  db.SavedResource
		interval  time.Duration
		immediate bool
	}
	acquireResourceCheckingLockReturns struct {
		result1 db.Lock
		result2 bool
		result3 error
	}
	AcquireResourceTypeCheckingLockStub        func(logger lager.Logger, resourceType db.SavedResourceType, interval time.Duration, immediate bool) (db.Lock, bool, error)
	acquireResourceTypeCheckingLockMutex       sync.RWMutex
	acquireResourceTypeCheckingLockArgsForCall []struct {
		logger       lager.Logger
		resourceType db.SavedResourceType
		interval     time.Duration
		immediate    bool
	}
	acquireResourceTypeCheckingLockReturns struct {
		result1 db.Lock
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRadarDB) GetPipelineName() string {
	fake.getPipelineNameMutex.Lock()
	fake.getPipelineNameArgsForCall = append(fake.getPipelineNameArgsForCall, struct{}{})
	fake.recordInvocation("GetPipelineName", []interface{}{})
	fake.getPipelineNameMutex.Unlock()
	if fake.GetPipelineNameStub != nil {
		return fake.GetPipelineNameStub()
	} else {
		return fake.getPipelineNameReturns.result1
	}
}

func (fake *FakeRadarDB) GetPipelineNameCallCount() int {
	fake.getPipelineNameMutex.RLock()
	defer fake.getPipelineNameMutex.RUnlock()
	return len(fake.getPipelineNameArgsForCall)
}

func (fake *FakeRadarDB) GetPipelineNameReturns(result1 string) {
	fake.GetPipelineNameStub = nil
	fake.getPipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRadarDB) GetPipelineID() int {
	fake.getPipelineIDMutex.Lock()
	fake.getPipelineIDArgsForCall = append(fake.getPipelineIDArgsForCall, struct{}{})
	fake.recordInvocation("GetPipelineID", []interface{}{})
	fake.getPipelineIDMutex.Unlock()
	if fake.GetPipelineIDStub != nil {
		return fake.GetPipelineIDStub()
	} else {
		return fake.getPipelineIDReturns.result1
	}
}

func (fake *FakeRadarDB) GetPipelineIDCallCount() int {
	fake.getPipelineIDMutex.RLock()
	defer fake.getPipelineIDMutex.RUnlock()
	return len(fake.getPipelineIDArgsForCall)
}

func (fake *FakeRadarDB) GetPipelineIDReturns(result1 int) {
	fake.GetPipelineIDStub = nil
	fake.getPipelineIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRadarDB) ScopedName(arg1 string) string {
	fake.scopedNameMutex.Lock()
	fake.scopedNameArgsForCall = append(fake.scopedNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ScopedName", []interface{}{arg1})
	fake.scopedNameMutex.Unlock()
	if fake.ScopedNameStub != nil {
		return fake.ScopedNameStub(arg1)
	} else {
		return fake.scopedNameReturns.result1
	}
}

func (fake *FakeRadarDB) ScopedNameCallCount() int {
	fake.scopedNameMutex.RLock()
	defer fake.scopedNameMutex.RUnlock()
	return len(fake.scopedNameArgsForCall)
}

func (fake *FakeRadarDB) ScopedNameArgsForCall(i int) string {
	fake.scopedNameMutex.RLock()
	defer fake.scopedNameMutex.RUnlock()
	return fake.scopedNameArgsForCall[i].arg1
}

func (fake *FakeRadarDB) ScopedNameReturns(result1 string) {
	fake.ScopedNameStub = nil
	fake.scopedNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRadarDB) TeamID() int {
	fake.teamIDMutex.Lock()
	fake.teamIDArgsForCall = append(fake.teamIDArgsForCall, struct{}{})
	fake.recordInvocation("TeamID", []interface{}{})
	fake.teamIDMutex.Unlock()
	if fake.TeamIDStub != nil {
		return fake.TeamIDStub()
	} else {
		return fake.teamIDReturns.result1
	}
}

func (fake *FakeRadarDB) TeamIDCallCount() int {
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	return len(fake.teamIDArgsForCall)
}

func (fake *FakeRadarDB) TeamIDReturns(result1 int) {
	fake.TeamIDStub = nil
	fake.teamIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRadarDB) Config() atc.Config {
	fake.configMutex.Lock()
	fake.configArgsForCall = append(fake.configArgsForCall, struct{}{})
	fake.recordInvocation("Config", []interface{}{})
	fake.configMutex.Unlock()
	if fake.ConfigStub != nil {
		return fake.ConfigStub()
	} else {
		return fake.configReturns.result1
	}
}

func (fake *FakeRadarDB) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *FakeRadarDB) ConfigReturns(result1 atc.Config) {
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 atc.Config
	}{result1}
}

func (fake *FakeRadarDB) IsPaused() (bool, error) {
	fake.isPausedMutex.Lock()
	fake.isPausedArgsForCall = append(fake.isPausedArgsForCall, struct{}{})
	fake.recordInvocation("IsPaused", []interface{}{})
	fake.isPausedMutex.Unlock()
	if fake.IsPausedStub != nil {
		return fake.IsPausedStub()
	} else {
		return fake.isPausedReturns.result1, fake.isPausedReturns.result2
	}
}

func (fake *FakeRadarDB) IsPausedCallCount() int {
	fake.isPausedMutex.RLock()
	defer fake.isPausedMutex.RUnlock()
	return len(fake.isPausedArgsForCall)
}

func (fake *FakeRadarDB) IsPausedReturns(result1 bool, result2 error) {
	fake.IsPausedStub = nil
	fake.isPausedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRadarDB) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct{}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	} else {
		return fake.reloadReturns.result1, fake.reloadReturns.result2
	}
}

func (fake *FakeRadarDB) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeRadarDB) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRadarDB) GetLatestVersionedResource(resourceName string) (db.SavedVersionedResource, bool, error) {
	fake.getLatestVersionedResourceMutex.Lock()
	fake.getLatestVersionedResourceArgsForCall = append(fake.getLatestVersionedResourceArgsForCall, struct {
		resourceName string
	}{resourceName})
	fake.recordInvocation("GetLatestVersionedResource", []interface{}{resourceName})
	fake.getLatestVersionedResourceMutex.Unlock()
	if fake.GetLatestVersionedResourceStub != nil {
		return fake.GetLatestVersionedResourceStub(resourceName)
	} else {
		return fake.getLatestVersionedResourceReturns.result1, fake.getLatestVersionedResourceReturns.result2, fake.getLatestVersionedResourceReturns.result3
	}
}

func (fake *FakeRadarDB) GetLatestVersionedResourceCallCount() int {
	fake.getLatestVersionedResourceMutex.RLock()
	defer fake.getLatestVersionedResourceMutex.RUnlock()
	return len(fake.getLatestVersionedResourceArgsForCall)
}

func (fake *FakeRadarDB) GetLatestVersionedResourceArgsForCall(i int) string {
	fake.getLatestVersionedResourceMutex.RLock()
	defer fake.getLatestVersionedResourceMutex.RUnlock()
	return fake.getLatestVersionedResourceArgsForCall[i].resourceName
}

func (fake *FakeRadarDB) GetLatestVersionedResourceReturns(result1 db.SavedVersionedResource, result2 bool, result3 error) {
	fake.GetLatestVersionedResourceStub = nil
	fake.getLatestVersionedResourceReturns = struct {
		result1 db.SavedVersionedResource
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRadarDB) GetResource(resourceName string) (db.SavedResource, bool, error) {
	fake.getResourceMutex.Lock()
	fake.getResourceArgsForCall = append(fake.getResourceArgsForCall, struct {
		resourceName string
	}{resourceName})
	fake.recordInvocation("GetResource", []interface{}{resourceName})
	fake.getResourceMutex.Unlock()
	if fake.GetResourceStub != nil {
		return fake.GetResourceStub(resourceName)
	} else {
		return fake.getResourceReturns.result1, fake.getResourceReturns.result2, fake.getResourceReturns.result3
	}
}

func (fake *FakeRadarDB) GetResourceCallCount() int {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return len(fake.getResourceArgsForCall)
}

func (fake *FakeRadarDB) GetResourceArgsForCall(i int) string {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return fake.getResourceArgsForCall[i].resourceName
}

func (fake *FakeRadarDB) GetResourceReturns(result1 db.SavedResource, result2 bool, result3 error) {
	fake.GetResourceStub = nil
	fake.getResourceReturns = struct {
		result1 db.SavedResource
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRadarDB) GetResourceType(resourceTypeName string) (db.SavedResourceType, bool, error) {
	fake.getResourceTypeMutex.Lock()
	fake.getResourceTypeArgsForCall = append(fake.getResourceTypeArgsForCall, struct {
		resourceTypeName string
	}{resourceTypeName})
	fake.recordInvocation("GetResourceType", []interface{}{resourceTypeName})
	fake.getResourceTypeMutex.Unlock()
	if fake.GetResourceTypeStub != nil {
		return fake.GetResourceTypeStub(resourceTypeName)
	} else {
		return fake.getResourceTypeReturns.result1, fake.getResourceTypeReturns.result2, fake.getResourceTypeReturns.result3
	}
}

func (fake *FakeRadarDB) GetResourceTypeCallCount() int {
	fake.getResourceTypeMutex.RLock()
	defer fake.getResourceTypeMutex.RUnlock()
	return len(fake.getResourceTypeArgsForCall)
}

func (fake *FakeRadarDB) GetResourceTypeArgsForCall(i int) string {
	fake.getResourceTypeMutex.RLock()
	defer fake.getResourceTypeMutex.RUnlock()
	return fake.getResourceTypeArgsForCall[i].resourceTypeName
}

func (fake *FakeRadarDB) GetResourceTypeReturns(result1 db.SavedResourceType, result2 bool, result3 error) {
	fake.GetResourceTypeStub = nil
	fake.getResourceTypeReturns = struct {
		result1 db.SavedResourceType
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRadarDB) PauseResource(resourceName string) error {
	fake.pauseResourceMutex.Lock()
	fake.pauseResourceArgsForCall = append(fake.pauseResourceArgsForCall, struct {
		resourceName string
	}{resourceName})
	fake.recordInvocation("PauseResource", []interface{}{resourceName})
	fake.pauseResourceMutex.Unlock()
	if fake.PauseResourceStub != nil {
		return fake.PauseResourceStub(resourceName)
	} else {
		return fake.pauseResourceReturns.result1
	}
}

func (fake *FakeRadarDB) PauseResourceCallCount() int {
	fake.pauseResourceMutex.RLock()
	defer fake.pauseResourceMutex.RUnlock()
	return len(fake.pauseResourceArgsForCall)
}

func (fake *FakeRadarDB) PauseResourceArgsForCall(i int) string {
	fake.pauseResourceMutex.RLock()
	defer fake.pauseResourceMutex.RUnlock()
	return fake.pauseResourceArgsForCall[i].resourceName
}

func (fake *FakeRadarDB) PauseResourceReturns(result1 error) {
	fake.PauseResourceStub = nil
	fake.pauseResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) UnpauseResource(resourceName string) error {
	fake.unpauseResourceMutex.Lock()
	fake.unpauseResourceArgsForCall = append(fake.unpauseResourceArgsForCall, struct {
		resourceName string
	}{resourceName})
	fake.recordInvocation("UnpauseResource", []interface{}{resourceName})
	fake.unpauseResourceMutex.Unlock()
	if fake.UnpauseResourceStub != nil {
		return fake.UnpauseResourceStub(resourceName)
	} else {
		return fake.unpauseResourceReturns.result1
	}
}

func (fake *FakeRadarDB) UnpauseResourceCallCount() int {
	fake.unpauseResourceMutex.RLock()
	defer fake.unpauseResourceMutex.RUnlock()
	return len(fake.unpauseResourceArgsForCall)
}

func (fake *FakeRadarDB) UnpauseResourceArgsForCall(i int) string {
	fake.unpauseResourceMutex.RLock()
	defer fake.unpauseResourceMutex.RUnlock()
	return fake.unpauseResourceArgsForCall[i].resourceName
}

func (fake *FakeRadarDB) UnpauseResourceReturns(result1 error) {
	fake.UnpauseResourceStub = nil
	fake.unpauseResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) SaveResourceVersions(arg1 atc.ResourceConfig, arg2 []atc.Version) error {
	var arg2Copy []atc.Version
	if arg2 != nil {
		arg2Copy = make([]atc.Version, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.saveResourceVersionsMutex.Lock()
	fake.saveResourceVersionsArgsForCall = append(fake.saveResourceVersionsArgsForCall, struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}{arg1, arg2Copy})
	fake.recordInvocation("SaveResourceVersions", []interface{}{arg1, arg2Copy})
	fake.saveResourceVersionsMutex.Unlock()
	if fake.SaveResourceVersionsStub != nil {
		return fake.SaveResourceVersionsStub(arg1, arg2)
	} else {
		return fake.saveResourceVersionsReturns.result1
	}
}

func (fake *FakeRadarDB) SaveResourceVersionsCallCount() int {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return len(fake.saveResourceVersionsArgsForCall)
}

func (fake *FakeRadarDB) SaveResourceVersionsArgsForCall(i int) (atc.ResourceConfig, []atc.Version) {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return fake.saveResourceVersionsArgsForCall[i].arg1, fake.saveResourceVersionsArgsForCall[i].arg2
}

func (fake *FakeRadarDB) SaveResourceVersionsReturns(result1 error) {
	fake.SaveResourceVersionsStub = nil
	fake.saveResourceVersionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) SaveResourceTypeVersion(arg1 atc.ResourceType, arg2 atc.Version) error {
	fake.saveResourceTypeVersionMutex.Lock()
	fake.saveResourceTypeVersionArgsForCall = append(fake.saveResourceTypeVersionArgsForCall, struct {
		arg1 atc.ResourceType
		arg2 atc.Version
	}{arg1, arg2})
	fake.recordInvocation("SaveResourceTypeVersion", []interface{}{arg1, arg2})
	fake.saveResourceTypeVersionMutex.Unlock()
	if fake.SaveResourceTypeVersionStub != nil {
		return fake.SaveResourceTypeVersionStub(arg1, arg2)
	} else {
		return fake.saveResourceTypeVersionReturns.result1
	}
}

func (fake *FakeRadarDB) SaveResourceTypeVersionCallCount() int {
	fake.saveResourceTypeVersionMutex.RLock()
	defer fake.saveResourceTypeVersionMutex.RUnlock()
	return len(fake.saveResourceTypeVersionArgsForCall)
}

func (fake *FakeRadarDB) SaveResourceTypeVersionArgsForCall(i int) (atc.ResourceType, atc.Version) {
	fake.saveResourceTypeVersionMutex.RLock()
	defer fake.saveResourceTypeVersionMutex.RUnlock()
	return fake.saveResourceTypeVersionArgsForCall[i].arg1, fake.saveResourceTypeVersionArgsForCall[i].arg2
}

func (fake *FakeRadarDB) SaveResourceTypeVersionReturns(result1 error) {
	fake.SaveResourceTypeVersionStub = nil
	fake.saveResourceTypeVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) SetResourceCheckError(resource db.SavedResource, err error) error {
	fake.setResourceCheckErrorMutex.Lock()
	fake.setResourceCheckErrorArgsForCall = append(fake.setResourceCheckErrorArgsForCall, struct {
		resource db.SavedResource
		err      error
	}{resource, err})
	fake.recordInvocation("SetResourceCheckError", []interface{}{resource, err})
	fake.setResourceCheckErrorMutex.Unlock()
	if fake.SetResourceCheckErrorStub != nil {
		return fake.SetResourceCheckErrorStub(resource, err)
	} else {
		return fake.setResourceCheckErrorReturns.result1
	}
}

func (fake *FakeRadarDB) SetResourceCheckErrorCallCount() int {
	fake.setResourceCheckErrorMutex.RLock()
	defer fake.setResourceCheckErrorMutex.RUnlock()
	return len(fake.setResourceCheckErrorArgsForCall)
}

func (fake *FakeRadarDB) SetResourceCheckErrorArgsForCall(i int) (db.SavedResource, error) {
	fake.setResourceCheckErrorMutex.RLock()
	defer fake.setResourceCheckErrorMutex.RUnlock()
	return fake.setResourceCheckErrorArgsForCall[i].resource, fake.setResourceCheckErrorArgsForCall[i].err
}

func (fake *FakeRadarDB) SetResourceCheckErrorReturns(result1 error) {
	fake.SetResourceCheckErrorStub = nil
	fake.setResourceCheckErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRadarDB) AcquireResourceCheckingLock(logger lager.Logger, resource db.SavedResource, interval time.Duration, immediate bool) (db.Lock, bool, error) {
	fake.acquireResourceCheckingLockMutex.Lock()
	fake.acquireResourceCheckingLockArgsForCall = append(fake.acquireResourceCheckingLockArgsForCall, struct {
		logger    lager.Logger
		resource  db.SavedResource
		interval  time.Duration
		immediate bool
	}{logger, resource, interval, immediate})
	fake.recordInvocation("AcquireResourceCheckingLock", []interface{}{logger, resource, interval, immediate})
	fake.acquireResourceCheckingLockMutex.Unlock()
	if fake.AcquireResourceCheckingLockStub != nil {
		return fake.AcquireResourceCheckingLockStub(logger, resource, interval, immediate)
	} else {
		return fake.acquireResourceCheckingLockReturns.result1, fake.acquireResourceCheckingLockReturns.result2, fake.acquireResourceCheckingLockReturns.result3
	}
}

func (fake *FakeRadarDB) AcquireResourceCheckingLockCallCount() int {
	fake.acquireResourceCheckingLockMutex.RLock()
	defer fake.acquireResourceCheckingLockMutex.RUnlock()
	return len(fake.acquireResourceCheckingLockArgsForCall)
}

func (fake *FakeRadarDB) AcquireResourceCheckingLockArgsForCall(i int) (lager.Logger, db.SavedResource, time.Duration, bool) {
	fake.acquireResourceCheckingLockMutex.RLock()
	defer fake.acquireResourceCheckingLockMutex.RUnlock()
	return fake.acquireResourceCheckingLockArgsForCall[i].logger, fake.acquireResourceCheckingLockArgsForCall[i].resource, fake.acquireResourceCheckingLockArgsForCall[i].interval, fake.acquireResourceCheckingLockArgsForCall[i].immediate
}

func (fake *FakeRadarDB) AcquireResourceCheckingLockReturns(result1 db.Lock, result2 bool, result3 error) {
	fake.AcquireResourceCheckingLockStub = nil
	fake.acquireResourceCheckingLockReturns = struct {
		result1 db.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRadarDB) AcquireResourceTypeCheckingLock(logger lager.Logger, resourceType db.SavedResourceType, interval time.Duration, immediate bool) (db.Lock, bool, error) {
	fake.acquireResourceTypeCheckingLockMutex.Lock()
	fake.acquireResourceTypeCheckingLockArgsForCall = append(fake.acquireResourceTypeCheckingLockArgsForCall, struct {
		logger       lager.Logger
		resourceType db.SavedResourceType
		interval     time.Duration
		immediate    bool
	}{logger, resourceType, interval, immediate})
	fake.recordInvocation("AcquireResourceTypeCheckingLock", []interface{}{logger, resourceType, interval, immediate})
	fake.acquireResourceTypeCheckingLockMutex.Unlock()
	if fake.AcquireResourceTypeCheckingLockStub != nil {
		return fake.AcquireResourceTypeCheckingLockStub(logger, resourceType, interval, immediate)
	} else {
		return fake.acquireResourceTypeCheckingLockReturns.result1, fake.acquireResourceTypeCheckingLockReturns.result2, fake.acquireResourceTypeCheckingLockReturns.result3
	}
}

func (fake *FakeRadarDB) AcquireResourceTypeCheckingLockCallCount() int {
	fake.acquireResourceTypeCheckingLockMutex.RLock()
	defer fake.acquireResourceTypeCheckingLockMutex.RUnlock()
	return len(fake.acquireResourceTypeCheckingLockArgsForCall)
}

func (fake *FakeRadarDB) AcquireResourceTypeCheckingLockArgsForCall(i int) (lager.Logger, db.SavedResourceType, time.Duration, bool) {
	fake.acquireResourceTypeCheckingLockMutex.RLock()
	defer fake.acquireResourceTypeCheckingLockMutex.RUnlock()
	return fake.acquireResourceTypeCheckingLockArgsForCall[i].logger, fake.acquireResourceTypeCheckingLockArgsForCall[i].resourceType, fake.acquireResourceTypeCheckingLockArgsForCall[i].interval, fake.acquireResourceTypeCheckingLockArgsForCall[i].immediate
}

func (fake *FakeRadarDB) AcquireResourceTypeCheckingLockReturns(result1 db.Lock, result2 bool, result3 error) {
	fake.AcquireResourceTypeCheckingLockStub = nil
	fake.acquireResourceTypeCheckingLockReturns = struct {
		result1 db.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRadarDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPipelineNameMutex.RLock()
	defer fake.getPipelineNameMutex.RUnlock()
	fake.getPipelineIDMutex.RLock()
	defer fake.getPipelineIDMutex.RUnlock()
	fake.scopedNameMutex.RLock()
	defer fake.scopedNameMutex.RUnlock()
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.isPausedMutex.RLock()
	defer fake.isPausedMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.getLatestVersionedResourceMutex.RLock()
	defer fake.getLatestVersionedResourceMutex.RUnlock()
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	fake.getResourceTypeMutex.RLock()
	defer fake.getResourceTypeMutex.RUnlock()
	fake.pauseResourceMutex.RLock()
	defer fake.pauseResourceMutex.RUnlock()
	fake.unpauseResourceMutex.RLock()
	defer fake.unpauseResourceMutex.RUnlock()
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	fake.saveResourceTypeVersionMutex.RLock()
	defer fake.saveResourceTypeVersionMutex.RUnlock()
	fake.setResourceCheckErrorMutex.RLock()
	defer fake.setResourceCheckErrorMutex.RUnlock()
	fake.acquireResourceCheckingLockMutex.RLock()
	defer fake.acquireResourceCheckingLockMutex.RUnlock()
	fake.acquireResourceTypeCheckingLockMutex.RLock()
	defer fake.acquireResourceTypeCheckingLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRadarDB) recordInvocation(key string, args []interface{}) {
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

var _ radar.RadarDB = new(FakeRadarDB)