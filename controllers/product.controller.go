package controllers

import (
	"net/http"

	"github.com/mhdianrush/go-token-auth-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]any{
		{
			"id":           1,
			"product_name": "Iphone 13",
			"stock":        100,
		},
		{
			"id":           2,
			"product_name": "Iphone 13 Pro",
			"stock":        50,
		},
		{
			"id":           3,
			"product_name": "Iphone 13 Pro Max",
			"stock":        10,
		},
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
