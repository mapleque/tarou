package client

import (
	"os"
	"text/template"
)

type Display struct {
	tpl  *template.Template
	data interface{}
	ops  interface{}
}

func (d *Display) SetTemplate(tpl string) error {
	var err error
	d.tpl, err = template.New("default").Parse(tpl)
	return err
}

func (d *Display) SetData(data interface{}) {
	d.data = data
}

func (d *Display) SetOps(ops interface{}) {
	d.ops = ops
}

func (d *Display) Output() {
	_ = d.tpl.Execute(os.Stdout, struct {
		Data interface{}
		Ops  interface{}
	}{
		Data: d.data,
		Ops:  d.ops,
	})
}
