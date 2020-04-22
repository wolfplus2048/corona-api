package corona

import (
	"context"
	"log"
	"plugin"
	"runtime"
)
var (
	app App
	aux Auxer
	pluginName string
)
func Default(name ...string) App {
	if app != nil {
		return app
	}
	if name != nil {
		pluginName = name[0]
	} else {
		pluginName = "corona-" + runtime.GOOS + ".so"
	}
	p, err := plugin.Open(pluginName)
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
	return app
}
func Aux() Auxer {
	if aux != nil {
		return aux
	}
	p, err := plugin.Open(pluginName)
	if err != nil {
		panic(err)
	}
	sym, err := p.Lookup("CoronaAux")
	if err != nil {
		panic(err)
	}
	a, ok := sym.(Auxer)
	if !ok {
		panic("expecting corona.auxer interface")
	}
	aux = a
	return aux
}
func GetSessionFromCtx(ctx context.Context) Session {
	sessionVal := ctx.Value("session")
	if sessionVal == nil {
		log.Print("ctx doesn't contain a session, are you calling GetSessionFromCtx from inside a remote?")
		return nil
	}
	return sessionVal.(Session)
}