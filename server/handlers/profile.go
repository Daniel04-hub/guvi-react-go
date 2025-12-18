package handlers

import (
	"context"
	"encoding/json"
	"guvi-project/db"
	"guvi-project/middleware"
	"guvi-project/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	// 1. Get email from context (set by middleware)
	email, ok := r.Context().Value(middleware.UserKey).(string)
	if !ok {
		http.Error(w, "Server Error: user context missing", http.StatusInternalServerError)
		return
	}

	// 2. Fetch from MongoDB
	coll := db.MongoClient.Database("guvi_db").Collection("profiles")
	var profile models.Profile
	err := coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&profile)
	if err != nil {
		// Return basic info if no profile doc exists yet
		profile = models.Profile{Email: email}
	}

	json.NewEncoder(w).Encode(profile)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value(middleware.UserKey).(string)
	if !ok {
		http.Error(w, "Server Error: user context missing", http.StatusInternalServerError)
		return
	}

	var req models.Profile
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Force email to match token
	req.Email = email

	// Upsert to MongoDB
	coll := db.MongoClient.Database("guvi_db").Collection("profiles")
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"email": email}
	update := bson.M{"$set": req}

	_, err := coll.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}
