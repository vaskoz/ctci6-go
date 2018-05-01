package one

import (
	"errors"
)

// Mistake demonstrates how to overflow a uint during a loop.
// pass true to demonstrate the overflow and return an error.
// pass false to not overflow and return nil instead of an error.
func Mistake(overflow bool) error {
	var i uint
	var maxUint = ^uint(0) // get max uint for this architecture
	var target uint = 1
	if overflow {
		target = 0
	}
	for i = 10; i >= target; i-- {
		if i == maxUint {
			return errors.New("equal to MaxUint")
		}
	}
	return nil
}
