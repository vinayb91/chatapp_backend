package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/vinayb91/chatapp_backend/config"
	"github.com/vinayb91/chatapp_backend/models"
	"github.com/vinayb91/chatapp_backend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	collection := config.GetCollection("users")
	filter := bson.M{"username": input.Username}
	err := collection.FindOne(context.TODO(), filter).Err()
	if err == nil {
		http.Error(w, "Username already exists", http.StatusUnauthorized)
		return
	}

	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)
	input.ID = primitive.NewObjectID()

	_, err = collection.InsertOne(context.TODO(), input)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"id":       input.ID.Hex(),
		"username": input.Username,
		"email":    input.Email,
	}

	tokenString, _ := utils.GenerateJWT(input.ID.Hex())
	utils.SetCookie(w, tokenString)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
