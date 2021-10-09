package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json: "name"`
	Username string `json: "username"`
	Age      int    `json: "age"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

type Post struct {
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}

func (user *User) hashpassword() {
	user.Password = fmt.Sprintf("%x", Encrypt([]byte((*user).Password), "password"))
}

func (user *User) validatePassword(loginPassword string) bool {
	return string(Decrypt([]byte((*user).Password), "password")) == loginPassword
}

func JSONError(res http.ResponseWriter, errVal map[string]string, code int) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.WriteHeader(code)

	jsonStr, _ := json.Marshal(errVal)
	res.Write(jsonStr)
}
