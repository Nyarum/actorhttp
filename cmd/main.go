package main

import (
	"net/http"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/actorhttp"
	"github.com/go-playground/pure"
	mw "github.com/go-playground/pure/_examples/middleware/logging-recovery"
)

type HandlerActor struct{}

func (state *HandlerActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case actorhttp.HTTPMessage:
		httpMessage := context.Message().(actorhttp.HTTPMessage)

		httpMessage.Response.Write([]byte("Hello World"))
	}
}

func main() {
	var (
		hf = actorhttp.New()
		p  = pure.New()
	)
	p.Use(mw.LoggingAndRecovery(true))

	p.Get("/", hf.ProtoHandler(&HandlerActor{}))

	http.ListenAndServe(":3007", p.Serve())
}
