package engin

import (
	"math/rand"
	"time"
)

// Unit is a unit in game which will growth.
type Unit struct {
	maxValue int64
	value    int64
}

// NewUnit create a new unit with init and max value.
func NewUnit(initValue, maxValue int64) *Unit {
	return &Unit{
		maxValue: maxValue,
		value:    initValue,
	}
}

// GrowUp the unit grow with value.
// Returns the value after grow up.
func (u *Unit) GrowUp(value int64) int64 {
	u.value += value
	return u.value
}

// AutoGrowUp the unit value grow up with interval d.
func (u *Unit) AutoGrowUp(d time.Duration, from, to int64) {
	go func() {
		for range time.NewTicker(d).C {
			_ = u.GrowUp(from + rand.Int63n(to-from))
		}
	}()
}

func (u *Unit) Value() int64 {
	return u.value
}
