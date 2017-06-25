# Actor HTTP

Inspired by github.com/AsynkronIT/protoactor-go

## Features

- Pool of actors
- State machine for handlers
- Supervisor for handlers
- Flexible with any HTTP router or framework

### Drawbacks
- Handle of request slower for 10-15 ns

## Documentation

### Install

- Soon...

## Other

### Benchmark

- Soon...

### Example

```
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
```

or you can find it in `cmd` folder