package corona

import (
	"context"
	"time"
)

type Session interface {
	Push(route string, v interface{}) error
	UID() string
	Bind(ctx context.Context, uid string)
	Kick(ctx context.Context) error
	OnClose(c func()) error
	Close()
	PeerAddr() string
	Set(key string, value interface{}) error
	Get(key string) interface{}
}

type App interface {
	AddAcceptor(addr string, certs ...string)
	Register(c Component, name string)
	RegisterRemote(c Component, name string)
	RegisterModule(module Module, name string) error
	Configure(isFrontend bool, serverType string, serverMetadata map[string]string)
	Start()
	NewCountTimer(interval time.Duration, count int, fn func()) int64
	RemoveTimer(id int64)
	AsyncTask(routine func()(interface{}, error), callback func(interface{}, error))
}

// Module is the interface that represent a module.
type Module interface {
	Init() error
	AfterInit()
	BeforeShutdown()
	Shutdown() error
}

// BindingStorage interface
type BindingStorage interface {
	GetUserFrontendID(uid, frontendType string) (string, error)
	PutBinding(uid string) error
}

// Component is the interface that represent a component.
type Component interface {
	Init()
	AfterInit()
	BeforeShutdown()
	Shutdown()
}
