package corona

import (
	"plugin"
)

func Default() App {
	p, err := plugin.Open("./corona.so")
	if err != nil {
		panic(err)
	}
	sym, err := p.Lookup("CoronaApp")
	if err != nil {
		panic(err)
	}
	app, ok := sym.(App)
	if !ok {
		panic("expecting corona.app interface")
	}
	return app
}