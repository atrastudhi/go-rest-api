package controllers

import (
	"net/http"
	"encoding/json"
	"strings"
	"github.com/atrastudhi/go-rest-api/models"
)

type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

func createErrorResponse(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	var res ErrorResponse
	res.Status = code
	res.Message = message

	var response, _ = json.Marshal(res)
	w.WriteHeader(code)
	w.Write(response)
}

func createSuccessResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	var res SuccessResponse
	res.Status = code
	res.Data = data

	var response, _ = json.Marshal(res)
	w.WriteHeader(code)
	w.Write(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	var err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		createErrorResponse(w, "internal server error", http.StatusInternalServerError)
		return
	} else if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Email) == "" {
		createErrorResponse(w, "all body is required", http.StatusBadRequest)
		return
	}

	createSuccessResponse(w, user, http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	users = append(users, models.User{Username: "badrun", Email: "me@badrun.com"})
	users = append(users, models.User{Username: "chikuy", Email: "me@chikuy.com"})
	users = append(users, models.User{Username: "jeci", Email: "me@jeci.com"})
	users = append(users, models.User{Username: "karatu", Email: "me@karatu.com"})

	createSuccessResponse(w, users, http.StatusOK)
}