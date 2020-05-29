package corona

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
	"log"
	"plugin"
	"runtime"
	"time"
)
var (
	app App
	aux Auxer
)
func Default(name ...string) App {
	if app != nil {
		return app
	}
	var pluginName string
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
	var ok bool
	app, ok = sym.(App)
	if !ok {
		panic("expecting corona.app interface")
	}
	sym, err = p.Lookup("CoronaAux")
	if err != nil {
		panic(err)
	}
	aux, ok = sym.(Auxer)
	if !ok {
		panic("expecting corona.auxer interface")
	}
	return app
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
func GetServerID() string {
	return aux.GetServerID()
}
func RPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message) error {
	return aux.RPC(ctx, routeStr, reply, arg)
}
func GetServersByType(t string) (map[string]*Server, error) {
	return aux.GetServersByType(t)
}