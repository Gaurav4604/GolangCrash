package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserRouteConnection(client *mongo.Client) {
	userCollection = client.Database("Appointy-Task").Collection("Users")
}

func sendUserResponse(res http.ResponseWriter, user User) {
	res.Header().Set("content-type", "application/json")
	user.Password = "" // to hide user's password
	json.NewEncoder(res).Encode(user)
}

func userInserthandler(res http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.hashpassword()
	result, _ := userCollection.InsertOne(ctx, user)

	user.ID = fmt.Sprint(result.InsertedID)
	sendUserResponse(res, user)
}

func getUserhandler(res http.ResponseWriter, r *http.Request, userID string) {
	var user User
	objectIDS, _ := primitive.ObjectIDFromHex(userID)
	err := userCollection.FindOne(ctx, bson.M{
		"_id": objectIDS,
	}).Decode(&user)
	if err != nil {
		JSONError(res, map[string]string{
			"Err": "User not found",
		}, 404)
	} else {
		println(user.ID)
		sendUserResponse(res, user)
	}
}

func MainUserRouter(res http.ResponseWriter, r *http.Request, endpointList []string) {
	switch len(endpointList) {
	case 1:
		if r.Method == "POST" {
			userInserthandler(res, r)
		} else {
			JSONError(res, map[string]string{
				"Error": "wrong http method used",
			}, 400)
		}
	case 2:
		if r.Method == "GET" {
			getUserhandler(res, r, endpointList[1])
		} else {
			JSONError(res, map[string]string{
				"Error": "wrong http method used",
			}, 400)
		}
	default:
		JSONError(res, map[string]string{
			"Error": "too many url params",
		}, 400)
	}
}
