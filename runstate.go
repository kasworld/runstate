// running state manage when multi gorutine is running
// 0 : runnging and can run all goroutine
// bit 0 : need stop all , somebody want stop
// bit 1 ~ 63 : use for each go routine
// can use max 63 goroutine
package runstate

import (
	"fmt"
	"sync/atomic"
)

type RunState uint64

func New() *RunState {
	return new(RunState)
}

func (rs RunState) String() string {
	return fmt.Sprintf("0b%0b", rs)
}

func (rs *RunState) SetBit(n uint) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(rs))
		newval := oldval | (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(rs), oldval, newval)
	}
}

func (rs *RunState) ClearBit(n uint) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(rs))
		newval := oldval &^ (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(rs), oldval, newval)
	}
}

func (rs *RunState) NegBit(n uint) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(rs))
		newval := oldval ^ (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(rs), oldval, newval)
	}
}

func (rs *RunState) GetBit(n uint) bool {
	val := atomic.LoadUint64((*uint64)(rs)) & (1 << n)
	return val != 0
}

//

// all clear
func (rs *RunState) CanRun() bool {
	return atomic.LoadUint64((*uint64)(rs)) == 0
}

// bit 0 set
func (rs *RunState) TryStop() {
	rs.SetBit(0)
}
