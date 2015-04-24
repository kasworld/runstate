// running state manage when multi gorutine is running
// 0 : runnging and can run all goroutine
// bit 0 : need stop all , somebody want stop
// bit 1 ~ 63 : use for each go routine
// can use max 63 goroutine
package runstate

import (
	"github.com/kasworld/bits64"
)

type RunState struct {
	bits64.Bits64
}

func New() *RunState {
	return new(RunState)
}

// all clear
func (rs *RunState) CanRun() bool {
	return rs.GetUint64() == 0
}

// bit 0 set
func (rs *RunState) TryStop() {
	rs.SetBit(0)
}
