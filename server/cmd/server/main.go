package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	opts := options.Client()

	opts.ApplyURI("mongodb://root:root@localhost:27017")

	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(fmt.Errorf("error connecting to MongoDB: %v", err))
	}

	db := c.Database("app")
	col := db.Collection("kv")

	handler := NewHandler(col)

	http.HandleFunc("POST /put", handler.handlePut)
	http.HandleFunc("POST /get", handler.handleGet)
	http.HandleFunc("POST /delete", handler.handleDelete)
	http.HandleFunc("POST /list", handler.handleList)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Errorf("server listening failed: %v", err))
	}
	fmt.Println("Connected to MongoDB")

}