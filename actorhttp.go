package actorhttp

import (
	"net/http"
	"sync"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type HTTPMessage struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type ActorHTTP struct {
	actorPools map[actor.Actor]*sync.Pool
}

func New() *ActorHTTP {
	return &ActorHTTP{
		actorPools: make(map[actor.Actor]*sync.Pool),
	}
}

func (ah *ActorHTTP) ProtoHandler(actorI actor.Actor) func(resp http.ResponseWriter, req *http.Request) {
	ap, ok := ah.actorPools[actorI]
	if !ok {
		ap = &sync.Pool{
			New: func() interface{} {
				props := actor.FromInstance(actorI)

				return actor.Spawn(props)
			},
		}

		ah.actorPools[actorI] = ap
	}

	return func(resp http.ResponseWriter, req *http.Request) {
		freeActor := ap.Get().(*actor.PID)
		freeActor.Tell(HTTPMessage{resp, req})
	}
}
