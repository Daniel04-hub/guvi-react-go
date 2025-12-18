package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"guvi-project/db"
	"guvi-project/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 1. Check if email exists
	var existsID int
	// Note: In real app, check error specifically for NoRows.
	err := db.SQLClient.QueryRow("SELECT id FROM users WHERE email = ?", req.Email).Scan(&existsID)
	if err == nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// 2. Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// 3. Insert User (Prepared Statement required)
	stmt, err := db.SQLClient.Prepare("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(req.Username, req.Email, string(hash)); err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 1. Fetch user by email
	var storedHash string
	var userID int
	err := db.SQLClient.QueryRow("SELECT id, password_hash FROM users WHERE email = ?", req.Email).Scan(&userID, &storedHash)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// 2. Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 3. Generate Token
	token := uuid.New().String()

	// 4. Store in Redis (Key: Token, Value: Email, TTL: 30m)
	// Storing email as session data. Could store userID JSON too.
	err = db.RedisClient.Set(context.Background(), token, req.Email, 30*time.Minute).Err()
	if err != nil {
		http.Error(w, "Redis error", http.StatusInternalServerError)
		return
	}

	// 5. Return Token
	json.NewEncoder(w).Encode(map[string]string{
		"token":   token,
		"message": "Login successful",
	})
}
