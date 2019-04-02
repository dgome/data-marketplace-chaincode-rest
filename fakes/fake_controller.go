// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	controller "github.com/lgsvl/data-marketplace-chaincode-rest/controller"
	resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
)

type FakeController struct {
	InvokeCreateResourceStub        func(resources.CreateResourceRequest) resources.CreateResourceResponse
	invokeCreateResourceMutex       sync.RWMutex
	invokeCreateResourceArgsForCall []struct {
		arg1 resources.CreateResourceRequest
	}
	invokeCreateResourceReturns struct {
		result1 resources.CreateResourceResponse
	}
	invokeCreateResourceReturnsOnCall map[int]struct {
		result1 resources.CreateResourceResponse
	}
	InvokeGetResourceStub        func(resources.GetResourceRequest) resources.GetResourceResponse
	invokeGetResourceMutex       sync.RWMutex
	invokeGetResourceArgsForCall []struct {
		arg1 resources.GetResourceRequest
	}
	invokeGetResourceReturns struct {
		result1 resources.GetResourceResponse
	}
	invokeGetResourceReturnsOnCall map[int]struct {
		result1 resources.GetResourceResponse
	}
	InvokeQueryStub        func(resources.QueryRequest) resources.QueryResponse
	invokeQueryMutex       sync.RWMutex
	invokeQueryArgsForCall []struct {
		arg1 resources.QueryRequest
	}
	invokeQueryReturns struct {
		result1 resources.QueryResponse
	}
	invokeQueryReturnsOnCall map[int]struct {
		result1 resources.QueryResponse
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeController) InvokeCreateResource(arg1 resources.CreateResourceRequest) resources.CreateResourceResponse {
	fake.invokeCreateResourceMutex.Lock()
	ret, specificReturn := fake.invokeCreateResourceReturnsOnCall[len(fake.invokeCreateResourceArgsForCall)]
	fake.invokeCreateResourceArgsForCall = append(fake.invokeCreateResourceArgsForCall, struct {
		arg1 resources.CreateResourceRequest
	}{arg1})
	fake.recordInvocation("InvokeCreateResource", []interface{}{arg1})
	fake.invokeCreateResourceMutex.Unlock()
	if fake.InvokeCreateResourceStub != nil {
		return fake.InvokeCreateResourceStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.invokeCreateResourceReturns
	return fakeReturns.result1
}

func (fake *FakeController) InvokeCreateResourceCallCount() int {
	fake.invokeCreateResourceMutex.RLock()
	defer fake.invokeCreateResourceMutex.RUnlock()
	return len(fake.invokeCreateResourceArgsForCall)
}

func (fake *FakeController) InvokeCreateResourceCalls(stub func(resources.CreateResourceRequest) resources.CreateResourceResponse) {
	fake.invokeCreateResourceMutex.Lock()
	defer fake.invokeCreateResourceMutex.Unlock()
	fake.InvokeCreateResourceStub = stub
}

func (fake *FakeController) InvokeCreateResourceArgsForCall(i int) resources.CreateResourceRequest {
	fake.invokeCreateResourceMutex.RLock()
	defer fake.invokeCreateResourceMutex.RUnlock()
	argsForCall := fake.invokeCreateResourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeController) InvokeCreateResourceReturns(result1 resources.CreateResourceResponse) {
	fake.invokeCreateResourceMutex.Lock()
	defer fake.invokeCreateResourceMutex.Unlock()
	fake.InvokeCreateResourceStub = nil
	fake.invokeCreateResourceReturns = struct {
		result1 resources.CreateResourceResponse
	}{result1}
}

func (fake *FakeController) InvokeCreateResourceReturnsOnCall(i int, result1 resources.CreateResourceResponse) {
	fake.invokeCreateResourceMutex.Lock()
	defer fake.invokeCreateResourceMutex.Unlock()
	fake.InvokeCreateResourceStub = nil
	if fake.invokeCreateResourceReturnsOnCall == nil {
		fake.invokeCreateResourceReturnsOnCall = make(map[int]struct {
			result1 resources.CreateResourceResponse
		})
	}
	fake.invokeCreateResourceReturnsOnCall[i] = struct {
		result1 resources.CreateResourceResponse
	}{result1}
}

func (fake *FakeController) InvokeGetResource(arg1 resources.GetResourceRequest) resources.GetResourceResponse {
	fake.invokeGetResourceMutex.Lock()
	ret, specificReturn := fake.invokeGetResourceReturnsOnCall[len(fake.invokeGetResourceArgsForCall)]
	fake.invokeGetResourceArgsForCall = append(fake.invokeGetResourceArgsForCall, struct {
		arg1 resources.GetResourceRequest
	}{arg1})
	fake.recordInvocation("InvokeGetResource", []interface{}{arg1})
	fake.invokeGetResourceMutex.Unlock()
	if fake.InvokeGetResourceStub != nil {
		return fake.InvokeGetResourceStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.invokeGetResourceReturns
	return fakeReturns.result1
}

func (fake *FakeController) InvokeGetResourceCallCount() int {
	fake.invokeGetResourceMutex.RLock()
	defer fake.invokeGetResourceMutex.RUnlock()
	return len(fake.invokeGetResourceArgsForCall)
}

func (fake *FakeController) InvokeGetResourceCalls(stub func(resources.GetResourceRequest) resources.GetResourceResponse) {
	fake.invokeGetResourceMutex.Lock()
	defer fake.invokeGetResourceMutex.Unlock()
	fake.InvokeGetResourceStub = stub
}

func (fake *FakeController) InvokeGetResourceArgsForCall(i int) resources.GetResourceRequest {
	fake.invokeGetResourceMutex.RLock()
	defer fake.invokeGetResourceMutex.RUnlock()
	argsForCall := fake.invokeGetResourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeController) InvokeGetResourceReturns(result1 resources.GetResourceResponse) {
	fake.invokeGetResourceMutex.Lock()
	defer fake.invokeGetResourceMutex.Unlock()
	fake.InvokeGetResourceStub = nil
	fake.invokeGetResourceReturns = struct {
		result1 resources.GetResourceResponse
	}{result1}
}

func (fake *FakeController) InvokeGetResourceReturnsOnCall(i int, result1 resources.GetResourceResponse) {
	fake.invokeGetResourceMutex.Lock()
	defer fake.invokeGetResourceMutex.Unlock()
	fake.InvokeGetResourceStub = nil
	if fake.invokeGetResourceReturnsOnCall == nil {
		fake.invokeGetResourceReturnsOnCall = make(map[int]struct {
			result1 resources.GetResourceResponse
		})
	}
	fake.invokeGetResourceReturnsOnCall[i] = struct {
		result1 resources.GetResourceResponse
	}{result1}
}

func (fake *FakeController) InvokeQuery(arg1 resources.QueryRequest) resources.QueryResponse {
	fake.invokeQueryMutex.Lock()
	ret, specificReturn := fake.invokeQueryReturnsOnCall[len(fake.invokeQueryArgsForCall)]
	fake.invokeQueryArgsForCall = append(fake.invokeQueryArgsForCall, struct {
		arg1 resources.QueryRequest
	}{arg1})
	fake.recordInvocation("InvokeQuery", []interface{}{arg1})
	fake.invokeQueryMutex.Unlock()
	if fake.InvokeQueryStub != nil {
		return fake.InvokeQueryStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.invokeQueryReturns
	return fakeReturns.result1
}

func (fake *FakeController) InvokeQueryCallCount() int {
	fake.invokeQueryMutex.RLock()
	defer fake.invokeQueryMutex.RUnlock()
	return len(fake.invokeQueryArgsForCall)
}

func (fake *FakeController) InvokeQueryCalls(stub func(resources.QueryRequest) resources.QueryResponse) {
	fake.invokeQueryMutex.Lock()
	defer fake.invokeQueryMutex.Unlock()
	fake.InvokeQueryStub = stub
}

func (fake *FakeController) InvokeQueryArgsForCall(i int) resources.QueryRequest {
	fake.invokeQueryMutex.RLock()
	defer fake.invokeQueryMutex.RUnlock()
	argsForCall := fake.invokeQueryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeController) InvokeQueryReturns(result1 resources.QueryResponse) {
	fake.invokeQueryMutex.Lock()
	defer fake.invokeQueryMutex.Unlock()
	fake.InvokeQueryStub = nil
	fake.invokeQueryReturns = struct {
		result1 resources.QueryResponse
	}{result1}
}

func (fake *FakeController) InvokeQueryReturnsOnCall(i int, result1 resources.QueryResponse) {
	fake.invokeQueryMutex.Lock()
	defer fake.invokeQueryMutex.Unlock()
	fake.InvokeQueryStub = nil
	if fake.invokeQueryReturnsOnCall == nil {
		fake.invokeQueryReturnsOnCall = make(map[int]struct {
			result1 resources.QueryResponse
		})
	}
	fake.invokeQueryReturnsOnCall[i] = struct {
		result1 resources.QueryResponse
	}{result1}
}

func (fake *FakeController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeCreateResourceMutex.RLock()
	defer fake.invokeCreateResourceMutex.RUnlock()
	fake.invokeGetResourceMutex.RLock()
	defer fake.invokeGetResourceMutex.RUnlock()
	fake.invokeQueryMutex.RLock()
	defer fake.invokeQueryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeController) recordInvocation(key string, args []interface{}) {
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

var _ controller.Controller = new(FakeController)