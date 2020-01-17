package main

import (
	"fmt"
	"net/http"

	"github.com/VictorOtavio/4okr/pkg/config"
	"github.com/gorilla/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/quarters", quarters)
	mux.HandleFunc("/objectives", objectives)
	mux.HandleFunc("/key-results", keyResults)

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", config.GetEnv("APP_PORT", "5000")),
		handlers.CompressHandler(mux),
	)
	if err != nil {
		panic("Erro ao iniciar o servidor HTTP: " + err.Error())
	}
}
