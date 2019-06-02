package handlers

import (
	"github.com/New-Era/servers/gateway/models/alerts"
	"github.com/New-Era/servers/gateway/models/devices"
	"github.com/New-Era/servers/gateway/sessions"
	webpush "github.com/SherClockHolmes/webpush-go"
)

// HandlerContext tracks the key that is used to sign and
// validate SessionIDs, the sessions.Store, and the users.Store
//Context holds contex values for multiple handler functions
type HandlerContext struct {
	SigningKey  string
	AlertStore  alerts.MySqlStore
	SessStore   sessions.RedisStore
	deviceStore devices.MongoStore
	Sockets     *SocketStore
	PubVapid    string
	PriVapid    string
}

//NewHandlerContext constructs a new HandlerContext,
//ensuring that the dependencies are valid values
func NewHandlerContext(signingKey string, alertStore *alerts.MySqlStore, sessStore *sessions.RedisStore, deviceStore *devices.MongoStore, connections *SocketStore) *HandlerContext {
	if signingKey == "" {
		panic("empty signing key")
	}
	if alertStore == nil {
		panic("nil alert")
	}
	if sessStore == nil {
		panic("nil session store")
	}
	if deviceStore == nil {
		panic("nil device store")
	}
	// pri/pub key for push notifications
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		panic("error generating VAPID Keys for Push Notifications")
	}

	return &HandlerContext{
		SigningKey:  signingKey,
		AlertStore:  *alertStore,
		SessStore:   *sessStore,
		deviceStore: *deviceStore,
		Sockets:     connections,
		PubVapid:    publicKey,
		PriVapid:    privateKey,
	}
}
