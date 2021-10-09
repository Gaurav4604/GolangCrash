package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection

func InitPostRouteConnection(client *mongo.Client) {
	postCollection = client.Database("Appointy-Task").Collection("Posts")
}

func setNewPageLimit(res http.ResponseWriter, value string) {
	expire := time.Now().Add(10 * time.Minute)
	cookie := http.Cookie{Name: "Page-Reached", Value: value, Path: "/", Expires: expire, MaxAge: 90000}
	http.SetCookie(res, &cookie)
}

func validateUser(userID string) bool {
	var user User
	objectID, _ := primitive.ObjectIDFromHex(userID)
	err := userCollection.FindOne(ctx, bson.M{
		"_id": objectID,
	}).Decode(&user)

	return err == nil
}

func getParticularPost(res http.ResponseWriter, r *http.Request, postID string) {
	var post Post
	objectIDS, _ := primitive.ObjectIDFromHex(postID)
	err := postCollection.FindOne(ctx, bson.M{
		"_id": objectIDS,
	}).Decode(&post)
	if err != nil {
		JSONError(res, map[string]string{
			"Err": "Post not found",
		}, 404)
	} else {
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(post)
	}
}

func postInsertHandler(res http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	if validateUser(post.UserID) {
		post.Timestamp = fmt.Sprintf("%d", time.Now().Unix()) // put time as unix timestamp
		postCollection.InsertOne(ctx, post)
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(post)
	} else {
		JSONError(res, map[string]string{
			"Err": "User not found",
		}, 404)
	}
}

func getAllPosts(res http.ResponseWriter, r *http.Request, userID string) {
	// val, _ := strconv.ParseInt(r.URL.Query().Get("param1"), 10, 64)

	var posts []Post
	cursor, err := postCollection.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		JSONError(res, map[string]string{
			"Err": "Not Found",
		}, 404)
	}
	if err = cursor.All(ctx, &posts); err != nil {
		JSONError(res, map[string]string{
			"Err": "Not Found",
		}, 404)
	}
	res.Header().Set("content-type", "application/json")

	// pagination via cookies and url params
	pages, _ := strconv.ParseInt(r.URL.Query().Get("pages"), 10, 64)
	if pages > 0 {
		for i, c := range r.Cookies() {
			if c.Name == "Page-Reached" {
				currentPage, err := strconv.ParseInt(c.Value, 10, 64)
				if err != nil {
					currentPage = 0
				}
				if (currentPage + pages) < int64(len(posts)) {
					setNewPageLimit(res, fmt.Sprint(currentPage+pages))
					json.NewEncoder(res).Encode(posts[currentPage:(currentPage + pages)])
				} else {
					json.NewEncoder(res).Encode(posts[:currentPage])
				}
			} else if i == len(r.Cookies())-1 {
				setNewPageLimit(res, fmt.Sprint(pages))
				json.NewEncoder(res).Encode(posts[:pages])
			}
		}

	} else {
		json.NewEncoder(res).Encode(posts)
	}
}

func MainPostsRouter(res http.ResponseWriter, r *http.Request, endpointList []string) {
	switch len(endpointList) {
	case 1:
		if r.Method == "POST" {
			postInsertHandler(res, r)
		} else {
			JSONError(res, map[string]string{
				"Error": "wrong http method used",
			}, 400)
		}
	case 2:
		if r.Method == "GET" {
			getParticularPost(res, r, endpointList[1])
		} else {
			JSONError(res, map[string]string{
				"Error": "wrong http method used",
			}, 400)
		}
	case 3:
		if r.Method == "GET" {
			getAllPosts(res, r, endpointList[2])
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
