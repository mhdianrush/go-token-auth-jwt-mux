package helper

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func ResponseJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
