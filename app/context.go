package app

import (
	"github.com/oatsaysai/simple-go-http-server/db"
	log "github.com/oatsaysai/simple-go-http-server/log"
)

type Context struct {
	Config        *Config
	Logger        log.Logger
	DB            *db.DB
	RemoteAddress string
}

func (app *App) NewContext() *Context {
	return &Context{
		Config: app.Config,
		Logger: app.Logger,
		DB:     app.DB,
	}
}

func (ctx *Context) WithLogger(logger log.Logger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}
