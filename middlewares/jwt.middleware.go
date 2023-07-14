package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mhdianrush/go-token-auth-jwt-mux/config"
	"github.com/mhdianrush/go-token-auth-jwt-mux/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]any{
					"message": "unauthorized",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		// exist token
		tokenString := c.Value
		claims := &config.JWTClaim{}
		// parse jkt token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_Key, nil
		})
		if err != nil {
			switch err {
			case jwt.ErrSignatureInvalid:
				response := map[string]string{
					"message": "unauthorize",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ErrTokenExpired:
				response := map[string]string{
					"message": "unauthorize, token expired!",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{
					"message": "unauthorize",
				}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		if !token.Valid {
			response := map[string]string{
				"message": "unauthorized",
			}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
		next.ServeHTTP(w, r)
	})
}
