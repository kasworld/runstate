package runstate

import (
	"testing"
)

func TestString(t *testing.T) {
	a := New()
	// a.SetBit(32)
	println(a.String())
}
