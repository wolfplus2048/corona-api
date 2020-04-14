package corona

import (
	"plugin"
)

var (
	app App
)
func init()  {
	p, err := plugin.Open("./corona.so")
	if err != nil {
		panic(err)
	}
	sym, err := p.Lookup("CoronaApp")
	if err != nil {
		panic(err)
	}
	a, ok := sym.(App)
	if !ok {
		panic("expecting corona.app interface")
	}
	app = a
}