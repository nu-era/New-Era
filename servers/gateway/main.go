package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/New-Era/servers/gateway/handlers"
	"github.com/New-Era/servers/gateway/models/devices"
	mgo "gopkg.in/mgo.v2"
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
