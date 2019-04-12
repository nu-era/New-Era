package main

import (
	"encoding/json"
	"fmt"
	"github.com/New-Era/servers/gateway/handlers"
	"github.com/New-Era/servers/gateway/models/devices"
	mgo "gopkg.in/mgo.v2"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync/atomic"
)

// main entry point for the server
func main() {
	// Connection to HTTPS
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlscert := os.Getenv("TLSCERT")
	tlskey := os.Getenv("TLSKEY")
	if len(tlskey) == 0 && len(tlscert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}

	// MONGO DB CONNECTION
	// get the address of the MongoDB server from an environment variable
	mongoAddr := os.Getenv("MONGO_ADDR")
	//default to "localhost"
	if len(mongoAddr) == 0 {
		mongoAddr = "localhost"
	}
	// Dialing MongoDB server
	mongoSess, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatalf("error dialing mongo: %v", err)
	}

	// TODO: construct a new MongoStore, provide mongoSess as well as a
	// database and collection name to use (device maybe?)

	conn := handlers.NewConnections()
	mongoStore := devices.NewMongoStore(mongoSess)
	handlerCtx := handlers.NewHandlerContext(mongoStore, conn)

	// messagingAddr := reqEnv("MESSAGESADDR")
	// summaryAddr := reqEnv("SUMMARYADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/time", handlers.TimeHandler)
	mux.HandleFunc("/device", handlerCtx.DevicesHandler)

	fmt.Printf("server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, mux))
}

// Director handles the transport of requests to proper endpoints
type Director func(r *http.Request)

// CustomDirectorRR directs requests for services that have
// multiple servers via Round Robin technique
func CustomDirectorRR(targets []string, hc *handlers.HandlerContext) Director {
	if len(targets) == 1 {
		dest, _ := url.Parse(targets[0])
		return CustomDirector(dest, hc)
	}
	var i int32
	i = 0
	url, _ := url.Parse(targets[int(i)%len(targets)])
	atomic.AddInt32(&i, 1)
	dest := url
	return func(r *http.Request) {
		r.Header.Del("X-Device") // remove any previous user
		tmp := handlers.SessionState{}
		_, _ = s.GetState(r, hc.Key, hc.Session, &tmp)
		if tmp.Device.ID != 0 { // set if user exists
			j, err := json.Marshal(tmp.Device)
			if err != nil {
				fmt.Println(err)
				return
			}
			r.Header.Set("X-Device", string(j))
		}
		r.Header.Add("X-Forwarded-Host", r.Host)
		r.URL.Scheme = "http"
		r.URL.Host = dest.String()
		r.Host = dest.String()
	}
}

// CustomDirector directs requests to a specified server and modifies the request
// before being passed along
func CustomDirector(target *url.URL, hc *handlers.HandlerContext) Director {
	return func(r *http.Request) {
		r.Header.Del("X-device") // remove any previous user
		tmp := handlers.SessionState{}
		_, _ = s.GetState(r, hc.Key, hc.Session, &tmp)
		if tmp.Device.ID != 0 { // set if user exists
			j, err := json.Marshal(tmp.User)
			if err != nil {
				fmt.Println(err)
				return
			}
			r.Header.Set("X-Device", string(j))
		}
		r.Header.Add("X-Forwarded-Host", r.Host)
		r.URL.Scheme = "http"
		r.URL.Host = target.String()
		r.Host = target.String()
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
