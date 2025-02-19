// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"context"
	apimodels "github.com/keptn/go-utils/pkg/api/models"
	"github.com/keptn/keptn/shipyard-controller/common"
	"github.com/keptn/keptn/shipyard-controller/models"
	"sync"
)

// ISequenceDispatcherMock is a mock implementation of handler.ISequenceDispatcher.
//
// 	func TestSomethingThatUsesISequenceDispatcher(t *testing.T) {
//
// 		// make and configure a mocked handler.ISequenceDispatcher
// 		mockedISequenceDispatcher := &ISequenceDispatcherMock{
// 			AddFunc: func(queueItem models.QueueItem) error {
// 				panic("mock out the Add method")
// 			},
// 			RemoveFunc: func(eventScope apimodels.KeptnContextExtendedCEScope) error {
// 				panic("mock out the Remove method")
// 			},
// 			RunFunc: func(ctx context.Context, startSequenceFunc func(event apimodels.KeptnContextExtendedCE) error)  {
// 				panic("mock out the Run method")
// 			},
// 			StopFunc: func()  {
// 				panic("mock out the Stop method")
// 			},
// 		}
//
// 		// use mockedISequenceDispatcher in code that requires handler.ISequenceDispatcher
// 		// and then make assertions.
//
// 	}
type ISequenceDispatcherMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(queueItem models.QueueItem) error

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(eventScope models.EventScope) error

	// RunFunc mocks the Run method.
	RunFunc func(ctx context.Context, mode common.SDMode, startSequenceFunc func(event apimodels.KeptnContextExtendedCE) error)

	// StopFunc mocks the Stop method.
	StopFunc func()

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// QueueItem is the queueItem argument value.
			QueueItem models.QueueItem
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// EventScope is the eventScope argument value.
			EventScope models.EventScope
		}
		// Run holds details about calls to the Run method.
		Run []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// StartSequenceFunc is the startSequenceFunc argument value.
			StartSequenceFunc func(event apimodels.KeptnContextExtendedCE) error
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
	}
	lockAdd    sync.RWMutex
	lockRemove sync.RWMutex
	lockRun    sync.RWMutex
	lockStop   sync.RWMutex
}

// Add calls AddFunc.
func (mock *ISequenceDispatcherMock) Add(queueItem models.QueueItem) error {
	if mock.AddFunc == nil {
		panic("ISequenceDispatcherMock.AddFunc: method is nil but ISequenceDispatcher.Add was just called")
	}
	callInfo := struct {
		QueueItem models.QueueItem
	}{
		QueueItem: queueItem,
	}
	mock.lockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	mock.lockAdd.Unlock()
	return mock.AddFunc(queueItem)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedISequenceDispatcher.AddCalls())
func (mock *ISequenceDispatcherMock) AddCalls() []struct {
	QueueItem models.QueueItem
} {
	var calls []struct {
		QueueItem models.QueueItem
	}
	mock.lockAdd.RLock()
	calls = mock.calls.Add
	mock.lockAdd.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *ISequenceDispatcherMock) Remove(eventScope models.EventScope) error {
	if mock.RemoveFunc == nil {
		panic("ISequenceDispatcherMock.RemoveFunc: method is nil but ISequenceDispatcher.Remove was just called")
	}
	callInfo := struct {
		EventScope models.EventScope
	}{
		EventScope: eventScope,
	}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc(eventScope)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//     len(mockedISequenceDispatcher.RemoveCalls())
func (mock *ISequenceDispatcherMock) RemoveCalls() []struct {
	EventScope models.EventScope
} {
	var calls []struct {
		EventScope models.EventScope
	}
	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}

// Run calls RunFunc.
func (mock *ISequenceDispatcherMock) Run(ctx context.Context, mode common.SDMode, startSequenceFunc func(event apimodels.KeptnContextExtendedCE) error) {
	if mock.RunFunc == nil {
		panic("ISequenceDispatcherMock.RunFunc: method is nil but ISequenceDispatcher.Run was just called")
	}
	callInfo := struct {
		Ctx               context.Context
		StartSequenceFunc func(event apimodels.KeptnContextExtendedCE) error
	}{
		Ctx:               ctx,
		StartSequenceFunc: startSequenceFunc,
	}
	mock.lockRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	mock.lockRun.Unlock()
	mock.RunFunc(ctx, mode, startSequenceFunc)
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//     len(mockedISequenceDispatcher.RunCalls())
func (mock *ISequenceDispatcherMock) RunCalls() []struct {
	Ctx               context.Context
	StartSequenceFunc func(event apimodels.KeptnContextExtendedCE) error
} {
	var calls []struct {
		Ctx               context.Context
		StartSequenceFunc func(event apimodels.KeptnContextExtendedCE) error
	}
	mock.lockRun.RLock()
	calls = mock.calls.Run
	mock.lockRun.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *ISequenceDispatcherMock) Stop() {
	if mock.StopFunc == nil {
		panic("ISequenceDispatcherMock.StopFunc: method is nil but ISequenceDispatcher.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//     len(mockedISequenceDispatcher.StopCalls())
func (mock *ISequenceDispatcherMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}
