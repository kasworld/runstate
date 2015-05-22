// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
