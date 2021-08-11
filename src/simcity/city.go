package simcity

import (
	"time"

	"github.com/mapleque/tarou/src/engin"
)

const ShellTemplate = `
Wellcome to our simcity!
+---------------------------------------+
Total	Unit: {{.Data.TotalUnit}}	Value: {{.Data.TotalValue}}
{{- range $index, $element := .Data.Units }}
>	Seq: {{$index}}	Value: {{$element.Value}}
{{- end }}
+---------------------------------------+
Operator list:
{{- range $index, $element := .Ops.Ops }}
{{$index}}	{{$element}}
{{- end }}
Please choose an operator and Press Enter:
`

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

func (c *City) TotalUnit() int {
	return len(c.unitPool)
}

func (c *City) TotalValue() int64 {
	var sum int64
	for _, u := range c.unitPool {
		sum += u.Value()
	}
	return sum
}

func (c *City) Units() []*engin.Unit {
	return c.unitPool
}
