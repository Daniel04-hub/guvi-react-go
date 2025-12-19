package handlers

import (
	"context"
	"encoding/json"
	"guvi-project/db"
	"guvi-project/middleware"
	"guvi-project/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value(middleware.UserKey).(string)
	if !ok {
		http.Error(w, "Server Error: user context missing", http.StatusInternalServerError)
		return
	}

	coll := db.MongoClient.Database("guvi_db").Collection("profiles")
	var profile models.Profile
	err := coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&profile)
	if err != nil {
		profile = models.Profile{Email: email}
	}

	json.NewEncoder(w).Encode(profile)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value(middleware.UserKey).(string)
	if !ok {
		log.Println("Profile Update Failed: User context missing") // Added log
		http.Error(w, "Server Error: user context missing", http.StatusInternalServerError)
		return
	}

	var profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		log.Println("Profile Update Failed: JSON Decode Error:", err) // Added log
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	profile.Email = email

	filter := bson.M{"email": email}
	update := bson.M{"$set": profile}

	opts := options.Update().SetUpsert(true)
	coll := db.MongoClient.Database("guvi_db").Collection("profiles")

	_, err := coll.UpdateOne(context.Background(), filter, update, opts) // Changed context.TODO() back to context.Background()
	if err != nil {
		log.Println("Profile Update Failed: MongoDB Error:", err) // Added log
		http.Error(w, "Failed to update profile", http.StatusInternalServerError) // Changed error message
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}
