package client

import (
	"fmt"
)

type Operator struct {
	ops  map[string][]func()
	tips map[string]string
}

func (o *Operator) Scan() {
	for {
		var op string
		fmt.Scanln(&op)
		if handlers, exist := o.ops[op]; exist {
			for _, h := range handlers {
				h()
			}
		}
	}
}

func (o *Operator) Bind(op, tips string, handlers ...func()) {
	if o.ops == nil {
		o.ops = map[string][]func(){}
		o.tips = map[string]string{}
	}
	o.ops[op] = handlers
	o.tips[op] = tips
}

func (o *Operator) Output() {
	fmt.Println("--------")
	for op, tips := range o.tips {
		fmt.Println(op, tips)
	}
	fmt.Println("Please choose an op: ")
}
