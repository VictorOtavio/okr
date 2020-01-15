package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VictorOtavio/4okr/pkg/config"
	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/quarters", quarters)
	mux.HandleFunc("/objectives", objectives)
	mux.HandleFunc("/key-results", keyResults)

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", config.GetEnv("APP_PORT", "8000")),
		handlers.CompressHandler(mux),
	)

	if err != nil {
		log.Fatal("Erro ao iniciar o servidor HTTP: ", err)
		return
	}
}

func quarters(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	fmt.Fprintln(w, "[]")
}

func objectives(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	fmt.Fprintln(w, "[]")
}

func keyResults(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	fmt.Fprintln(w, "[]")
}
