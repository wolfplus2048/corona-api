package corona

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"plugin"
	"runtime"
	"time"
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
func Config() *viper.Viper {
	return aux.GetConfig()
}
func AsyncTask(routine func()(interface{}, error), callback func(interface{}, error))  {
	aux.AsyncTask(routine, callback)
}
func NewCountTimer(interval time.Duration, count int, fn func()) int64 {
	return aux.NewCountTimer(interval, count, fn)
}
func RemoveTimer(id int64) {
	aux.RemoveTimer(id)
}