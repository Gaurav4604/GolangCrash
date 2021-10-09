package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context

func rootRouteHandler(res http.ResponseWriter, r *http.Request) {
	endpointList := strings.Split(r.URL.Path[8:], "/")
	fmt.Println(r.Method)

	switch endpointList[0] {
	case "users":
		MainUserRouter(res, r, endpointList)
	}
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// the mongo server connection context works for 10 minutes then times out
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Minute)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	InitUserRouteConnection(client)

	http.HandleFunc("/api/v1/", rootRouteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
