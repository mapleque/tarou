package main

import (
	"github.com/mapleque/tarou/src/client"
	"github.com/mapleque/tarou/src/simcity"
)

func main() {
	c := client.New()
	city := simcity.New()

	_ = c.Display.SetTemplate(simcity.ShellTemplate)
	c.Display.SetData(city)

	c.Operator.Bind("n", "new build", city.CreateUnit)
	c.Operator.Bind("q", "exit", c.Exit)
	c.Display.SetOps(c.Operator)

	c.Refresh()

	c.Start()
}
