package client

import (
	"fmt"
)

type Operator struct {
	ops  map[string][]func()
	tips []string
}

func (o *Operator) Scan() {
	for {
		var op string
		fmt.Scanf("%1s", &op)
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
	}
	if _, exist := o.ops[op]; !exist {
		o.ops[op] = handlers
		o.tips = append(o.tips, fmt.Sprintf("%s: %s", op, tips))
	}

}

func (o *Operator) Ops() []string {
	return o.tips
}
