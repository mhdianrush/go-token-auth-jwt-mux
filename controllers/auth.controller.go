package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mhdianrush/go-token-auth-jwt-mux/config"
	"github.com/mhdianrush/go-token-auth-jwt-mux/entities"
	"github.com/mhdianrush/go-token-auth-jwt-mux/helper"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var logger = logrus.New()

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// to catch json input
	var userInput entities.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash password with bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Println(err)
	}
	userInput.Password = string(hashPassword)

	// insert to db
	if err = config.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{
		"message": "registration success",
	}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
