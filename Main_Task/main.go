package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx context.Context

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

type Post struct {
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}

func rootHandler(res http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Path[8:]
	fmt.Println(r.Method)
	res.Header().Set("content-type", "application/json")
	if endpoint != "" {
		cur, err := collection.Find(ctx, bson.M{})

		if err != nil {
			panic(err)
		}
		defer cur.Close(ctx)

		var posts []Post
		if err = cur.All(ctx, &posts); err != nil {
			panic(err)
		}
		json.NewEncoder(res).Encode(posts)
	}
}

func main() {
	ciphertext := Encrypt([]byte("Hello World"), "password")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := Decrypt(ciphertext, "password")
	fmt.Printf("Decrypted: %s\n", plaintext)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("Appointy-Task").Collection("Users")

	http.HandleFunc("/api/v1/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	defer client.Disconnect(ctx)
}
