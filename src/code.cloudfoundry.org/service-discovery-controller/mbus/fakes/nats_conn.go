// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/service-discovery-controller/mbus"
	nats "github.com/nats-io/go-nats"
)

type NatsConn struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	ConnectedUrlStub        func() string
	connectedUrlMutex       sync.RWMutex
	connectedUrlArgsForCall []struct {
	}
	connectedUrlReturns struct {
		result1 string
	}
	connectedUrlReturnsOnCall map[int]struct {
		result1 string
	}
	FlushStub        func() error
	flushMutex       sync.RWMutex
	flushArgsForCall []struct {
	}
	flushReturns struct {
		result1 error
	}
	flushReturnsOnCall map[int]struct {
		result1 error
	}
	PublishMsgStub        func(*nats.Msg) error
	publishMsgMutex       sync.RWMutex
	publishMsgArgsForCall []struct {
		arg1 *nats.Msg
	}
	publishMsgReturns struct {
		result1 error
	}
	publishMsgReturnsOnCall map[int]struct {
		result1 error
	}
	SubscribeStub        func(string, nats.MsgHandler) (*nats.Subscription, error)
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		arg1 string
		arg2 nats.MsgHandler
	}
	subscribeReturns struct {
		result1 *nats.Subscription
		result2 error
	}
	subscribeReturnsOnCall map[int]struct {
		result1 *nats.Subscription
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NatsConn) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *NatsConn) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *NatsConn) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *NatsConn) ConnectedUrl() string {
	fake.connectedUrlMutex.Lock()
	ret, specificReturn := fake.connectedUrlReturnsOnCall[len(fake.connectedUrlArgsForCall)]
	fake.connectedUrlArgsForCall = append(fake.connectedUrlArgsForCall, struct {
	}{})
	stub := fake.ConnectedUrlStub
	fakeReturns := fake.connectedUrlReturns
	fake.recordInvocation("ConnectedUrl", []interface{}{})
	fake.connectedUrlMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *NatsConn) ConnectedUrlCallCount() int {
	fake.connectedUrlMutex.RLock()
	defer fake.connectedUrlMutex.RUnlock()
	return len(fake.connectedUrlArgsForCall)
}

func (fake *NatsConn) ConnectedUrlCalls(stub func() string) {
	fake.connectedUrlMutex.Lock()
	defer fake.connectedUrlMutex.Unlock()
	fake.ConnectedUrlStub = stub
}

func (fake *NatsConn) ConnectedUrlReturns(result1 string) {
	fake.connectedUrlMutex.Lock()
	defer fake.connectedUrlMutex.Unlock()
	fake.ConnectedUrlStub = nil
	fake.connectedUrlReturns = struct {
		result1 string
	}{result1}
}

func (fake *NatsConn) ConnectedUrlReturnsOnCall(i int, result1 string) {
	fake.connectedUrlMutex.Lock()
	defer fake.connectedUrlMutex.Unlock()
	fake.ConnectedUrlStub = nil
	if fake.connectedUrlReturnsOnCall == nil {
		fake.connectedUrlReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.connectedUrlReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *NatsConn) Flush() error {
	fake.flushMutex.Lock()
	ret, specificReturn := fake.flushReturnsOnCall[len(fake.flushArgsForCall)]
	fake.flushArgsForCall = append(fake.flushArgsForCall, struct {
	}{})
	stub := fake.FlushStub
	fakeReturns := fake.flushReturns
	fake.recordInvocation("Flush", []interface{}{})
	fake.flushMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *NatsConn) FlushCallCount() int {
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	return len(fake.flushArgsForCall)
}

func (fake *NatsConn) FlushCalls(stub func() error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = stub
}

func (fake *NatsConn) FlushReturns(result1 error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = nil
	fake.flushReturns = struct {
		result1 error
	}{result1}
}

func (fake *NatsConn) FlushReturnsOnCall(i int, result1 error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = nil
	if fake.flushReturnsOnCall == nil {
		fake.flushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.flushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NatsConn) PublishMsg(arg1 *nats.Msg) error {
	fake.publishMsgMutex.Lock()
	ret, specificReturn := fake.publishMsgReturnsOnCall[len(fake.publishMsgArgsForCall)]
	fake.publishMsgArgsForCall = append(fake.publishMsgArgsForCall, struct {
		arg1 *nats.Msg
	}{arg1})
	stub := fake.PublishMsgStub
	fakeReturns := fake.publishMsgReturns
	fake.recordInvocation("PublishMsg", []interface{}{arg1})
	fake.publishMsgMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *NatsConn) PublishMsgCallCount() int {
	fake.publishMsgMutex.RLock()
	defer fake.publishMsgMutex.RUnlock()
	return len(fake.publishMsgArgsForCall)
}

func (fake *NatsConn) PublishMsgCalls(stub func(*nats.Msg) error) {
	fake.publishMsgMutex.Lock()
	defer fake.publishMsgMutex.Unlock()
	fake.PublishMsgStub = stub
}

func (fake *NatsConn) PublishMsgArgsForCall(i int) *nats.Msg {
	fake.publishMsgMutex.RLock()
	defer fake.publishMsgMutex.RUnlock()
	argsForCall := fake.publishMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *NatsConn) PublishMsgReturns(result1 error) {
	fake.publishMsgMutex.Lock()
	defer fake.publishMsgMutex.Unlock()
	fake.PublishMsgStub = nil
	fake.publishMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *NatsConn) PublishMsgReturnsOnCall(i int, result1 error) {
	fake.publishMsgMutex.Lock()
	defer fake.publishMsgMutex.Unlock()
	fake.PublishMsgStub = nil
	if fake.publishMsgReturnsOnCall == nil {
		fake.publishMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NatsConn) Subscribe(arg1 string, arg2 nats.MsgHandler) (*nats.Subscription, error) {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		arg1 string
		arg2 nats.MsgHandler
	}{arg1, arg2})
	stub := fake.SubscribeStub
	fakeReturns := fake.subscribeReturns
	fake.recordInvocation("Subscribe", []interface{}{arg1, arg2})
	fake.subscribeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *NatsConn) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *NatsConn) SubscribeCalls(stub func(string, nats.MsgHandler) (*nats.Subscription, error)) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = stub
}

func (fake *NatsConn) SubscribeArgsForCall(i int) (string, nats.MsgHandler) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	argsForCall := fake.subscribeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *NatsConn) SubscribeReturns(result1 *nats.Subscription, result2 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 *nats.Subscription
		result2 error
	}{result1, result2}
}

func (fake *NatsConn) SubscribeReturnsOnCall(i int, result1 *nats.Subscription, result2 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 *nats.Subscription
			result2 error
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 *nats.Subscription
		result2 error
	}{result1, result2}
}

func (fake *NatsConn) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.connectedUrlMutex.RLock()
	defer fake.connectedUrlMutex.RUnlock()
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	fake.publishMsgMutex.RLock()
	defer fake.publishMsgMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NatsConn) recordInvocation(key string, args []interface{}) {
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

var _ mbus.NatsConn = new(NatsConn)