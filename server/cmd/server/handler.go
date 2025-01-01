package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	coll *mongo.Collection
}

type Document struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

func NewHandler(coll *mongo.Collection) *Handler {
	return &Handler{coll: coll}
}

type PutReqBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PutRespBody struct {
	Ok bool `json:"ok"`
}

func (h *Handler) handlePut(w http.ResponseWriter, r *http.Request) {
	reqBody := PutReqBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to decode req body: %w", err).Error(), http.StatusBadRequest)
		return
	}

	doc := &Document{Key: reqBody.Key, Value: reqBody.Value}
	filter := bson.M{"key": doc.Key}
	update := bson.M{"$set": doc}
	opts := options.Update().SetUpsert(true)

	_, err = h.coll.UpdateOne(r.Context(), filter, update, opts)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to update document: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	respBody := PutRespBody{Ok: true}
	err = json.NewEncoder(w).Encode(respBody)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to encode resp body: %w", err).Error(), http.StatusInternalServerError)
	}
}

type GetReqBody struct {
	Key string `json:"key"`
}

type GetRespBody struct {
	Value string `json:"value"`
	Ok    bool   `json:"ok"`
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	reqBody := GetReqBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to decode req body: %w", err).Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"key": reqBody.Key}
	doc := &Document{}
	err = h.coll.FindOne(r.Context(), filter).Decode(doc)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to find document: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	respBody := GetRespBody{Value: doc.Value, Ok: true}
	err = json.NewEncoder(w).Encode(respBody)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to encode resp body: %w", err).Error(), http.StatusInternalServerError)
	}
}

type DeleteReqBody struct {
	Key string `json:"key"`
}

type DeleteRespBody struct {
	Ok bool `json:"ok"`
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	reqBody := DeleteReqBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to decode req body: %w", err).Error(), http.StatusBadRequest)
		return
	}

	res, err := h.coll.DeleteOne(r.Context(), bson.M{"key": reqBody.Key})
	if err != nil {
		http.Error(w, fmt.Errorf("failed to delete document: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	respBody := DeleteRespBody{Ok: res.DeletedCount > 0}
	err = json.NewEncoder(w).Encode(respBody)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to encode resp body: %w", err).Error(), http.StatusInternalServerError)
	}
}

type ListReqBody struct{}

type ListRespBody struct {
	Docs []Document `json:"items"`
}

func (h *Handler) handleList(w http.ResponseWriter, r *http.Request) {
	cur, err := h.coll.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, fmt.Errorf("failed to find documents: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	docs := []Document{}
	err = cur.All(r.Context(), &docs)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to decode documents: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	respBody := ListRespBody{Docs: docs}
	err = json.NewEncoder(w).Encode(respBody)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to encode resp body: %w", err).Error(), http.StatusInternalServerError)
	}
}