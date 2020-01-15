package main

import (
	"fmt"
	"net/http"
)

type keyResult struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func keyResults(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	fmt.Fprintln(w, "[]")
}
