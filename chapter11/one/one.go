package one

import (
	"errors"
)

func Mistake() error {
	var i uint
	var maxUint uint = 0
	maxUint-- // get max uint for this architecture
	for i = 1; i >= 0; i-- {
		if i == maxUint {
			return errors.New("equal to MaxUint")
		}
	}
	return nil
}
