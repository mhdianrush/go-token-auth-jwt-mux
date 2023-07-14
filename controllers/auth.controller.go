package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mhdianrush/go-token-auth-jwt-mux/config"
	"github.com/mhdianrush/go-token-auth-jwt-mux/entities"
	"github.com/mhdianrush/go-token-auth-jwt-mux/helper"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logrus.New()

func Login(w http.ResponseWriter, r *http.Request) {
	// to catch user input
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

	// verify user data based on username
	var user entities.User

	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "username or password wrong",
			}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{
				"message": err.Error(),
			}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// verify user password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		response := map[string]string{
			"message": "username or password wrong",
		}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// generate token
	// token will expire after 1 minute
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-token-auth-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare algorithm for sign in
	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgorithm.SignedString(config.JWT_Key)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// if success, token will be set at cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{
		"message": "login success",
	}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// to catch user input
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
