/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * promise/future.go                                      *
 *                                                        *
 * future promise implementation for Go.                  *
 *                                                        *
 * LastModified: Aug 13, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package promise

import (
	"sync/atomic"
	"time"
)

type subscriber struct {
	onFulfilled OnFulfilled
	onRejected  OnRejected
	next        Promise
}

type future struct {
	value       interface{}
	reason      error
	state       uint32
	subscribers []subscriber
}

// New creates a PENDING Promise object
func New() Promise {
	return new(future)
}

func (p *future) Then(onFulfilled OnFulfilled, rest ...OnRejected) Promise {
	var onRejected OnRejected
	if len(rest) > 0 {
		onRejected = rest[0]
	}
	next := New()
	switch State(p.state) {
	case FULFILLED:
		if onFulfilled == nil {
			return fulfilled{p.value}
		}
		resolve(next, onFulfilled, p.value)
	case REJECTED:
		if onRejected == nil {
			return rejected{p.reason}
		}
		reject(next, onRejected, p.reason)
	default:
		p.subscribers = append(p.subscribers,
			subscriber{onFulfilled, onRejected, next})
	}
	return next
}

func (p *future) Catch(onRejected OnRejected, test ...TestFunc) Promise {
	if len(test) == 0 || test[0] == nil {
		return p.Then(nil, onRejected)
	}
	return p.Then(nil, func(e error) (interface{}, error) {
		if test[0](e) {
			return p.Then(nil, onRejected), nil
		}
		return nil, e
	})
}

func (p *future) Complete(onCompleted OnCompleted) Promise {
	return p.Then(OnFulfilled(onCompleted), func(e error) (interface{}, error) {
		return onCompleted(e)
	})
}

func (p *future) WhenComplete(action func()) Promise {
	return p.Then(func(v interface{}) (interface{}, error) {
		action()
		return v, nil
	}, func(e error) (interface{}, error) {
		action()
		return nil, e
	})
}

func (p *future) Done(onFulfilled OnFulfilled, onRejected ...OnRejected) {
	p.
		Then(onFulfilled, onRejected...).
		Then(nil, func(e error) (interface{}, error) {
			go panic(e)
			return nil, nil
		})
}

func (p *future) State() State {
	return State(p.state)
}

func thenableCatch(promise Promise, done *uint32) {
	if e := recover(); e != nil && atomic.CompareAndSwapUint32(done, 0, 1) {
		promise.Reject(NewPanicError(e))
	}
}

func getThenableOnFullfilled(promise Promise, done *uint32) OnFulfilled {
	return func(y interface{}) (interface{}, error) {
		if atomic.CompareAndSwapUint32(done, 0, 1) {
			promise.Resolve(y)
		}
		return nil, nil
	}
}

func getThenableOnRejected(promise Promise, done *uint32) OnRejected {
	return func(e error) (interface{}, error) {
		if atomic.CompareAndSwapUint32(done, 0, 1) {
			promise.Reject(e)
		}
		return nil, nil
	}
}

func resolveThenable(p *future, thenable Thenable) {
	var done uint32
	defer thenableCatch(p, &done)
	thenable.Then(
		getThenableOnFullfilled(p, &done),
		getThenableOnRejected(p, &done))
}

func resloveValue(p *future, value interface{}) {
	if atomic.CompareAndSwapUint32(&p.state, uint32(PENDING), uint32(FULFILLED)) {
		p.value = value
		subscribers := p.subscribers
		p.subscribers = nil
		for _, subscriber := range subscribers {
			resolve(subscriber.next, subscriber.onFulfilled, value)
		}
	}
}

func (p *future) Resolve(value interface{}) {
	if promise, ok := value.(*future); ok && promise == p {
		p.Reject(TypeError{"Self resolution"})
		return
	}
	if promise, ok := value.(Promise); ok {
		promise.Fill(p)
		return
	}
	if thenable, ok := value.(Thenable); ok {
		resolveThenable(p, thenable)
		return
	}
	resloveValue(p, value)
}

func (p *future) Reject(reason error) {
	if atomic.CompareAndSwapUint32(&p.state, uint32(PENDING), uint32(REJECTED)) {
		p.reason = reason
		subscribers := p.subscribers
		p.subscribers = nil
		for _, subscriber := range subscribers {
			reject(subscriber.next, subscriber.onRejected, reason)
		}
	}
}

func (p *future) Fill(promise Promise) {
	resolveFunc := func(v interface{}) (interface{}, error) {
		promise.Resolve(v)
		return nil, nil
	}
	rejectFunc := func(e error) (interface{}, error) {
		promise.Reject(e)
		return nil, nil
	}
	p.Then(resolveFunc, rejectFunc)
}

func (p *future) Timeout(duration time.Duration, reason ...error) Promise {
	return timeout(p, duration, reason...)
}

func (p *future) Delay(duration time.Duration) Promise {
	next := New()
	p.Then(func(v interface{}) (interface{}, error) {
		go func() {
			time.Sleep(duration)
			next.Resolve(v)
		}()
		return nil, nil
	}, func(e error) (interface{}, error) {
		next.Reject(e)
		return nil, nil
	})
	return next
}

func (p *future) Tap(onfulfilledSideEffect OnfulfilledSideEffect) Promise {
	return tap(p, onfulfilledSideEffect)
}

func (p *future) Get() (interface{}, error) {
	c := make(chan interface{})
	p.Then(func(v interface{}) (interface{}, error) {
		c <- v
		return nil, nil
	}, func(e error) (interface{}, error) {
		c <- e
		return nil, nil
	})
	v := <-c
	if e, ok := v.(error); ok {
		return nil, e
	}
	return v, nil
}