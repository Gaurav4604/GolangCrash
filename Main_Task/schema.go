package main

import (
	"fmt"
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
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    string `json:"userID,omitempty" bson:"userID,omitempty"`
	Caption   string `bson:"caption, omitempty"`
	Body      string `bson:"body,omitempty"`
	ImageURL  string `"bson:"imageurl, omitempty`
	Timestamp string `"bson":timestamp, omitempty`
}

func (user *User) hashpassword() {
	user.Password = fmt.Sprintf("%x", Encrypt([]byte((*user).Password), "password"))
}

func (user *User) validatePassword(loginPassword string) bool {
	return string(Decrypt([]byte((*user).Password), "password")) == loginPassword
}
