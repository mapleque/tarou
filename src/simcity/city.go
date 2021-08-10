package simcity

import (
	"time"

	"github.com/mapleque/tarou/src/engin"
)

const ShellTemplate = "unit: {{.UnitNum}}, Value: {{.Value}}\n"

type City struct {
	unitPool []*engin.Unit
}

func New() *City {
	return &City{}
}

func (c *City) CreateUnit() {
	unit := engin.NewUnit(0, 4*365*24*60*60)
	unit.AutoGrowUp(time.Second, 1, 4)
	c.unitPool = append(c.unitPool, unit)
}

func (c *City) UnitNum() int {
	return len(c.unitPool)
}

func (c *City) Value() int64 {
	var sum int64
	for _, u := range c.unitPool {
		sum += u.Value()
	}
	return sum
}
