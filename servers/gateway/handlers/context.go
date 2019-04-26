package handlers

import (
	"github.com/New-Era/servers/gateway/models/devices"
	"github.com/New-Era/servers/gateway/sessions"
)

// HandlerContext tracks the key that is used to sign and
// validate SessionIDs, the sessions.Store, and the users.Store
//Context holds contex values for multiple handler functions
type HandlerContext struct {
	SigningKey  string
	SessStore   sessions.RedisStore
	deviceStore devices.MongoStore
	Sockets     *SocketStore
}

//NewHandlerContext constructs a new HandlerContext,
//ensuring that the dependencies are valid values
func NewHandlerContext(signingKey string, sessStore *sessions.RedisStore, deviceStore *devices.MongoStore, connections *SocketStore) *HandlerContext {
	if signingKey == "" {
		panic("empty signing key")
	}
	if sessStore == nil {
		panic("nil session store")
	}
	if deviceStore == nil {
		panic("nil device store")
	}
	return &HandlerContext{
		SigningKey:  signingKey,
		SessStore:   *sessStore,
		deviceStore: *deviceStore,
		Sockets:     connections,
	}
}
