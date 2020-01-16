package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorOtavio/4okr/pkg/config"
)

type quarter struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func quarters(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	db, _ := sql.Open("sqlite3", config.GetEnv("DB_PATH", "../database/database.sqlite"))
	rows, _ := db.Query("SELECT id, title, start_date, end_date FROM quarters")

	quarters := make([]quarter, 0)

	for rows.Next() {
		var qrt quarter
		err := rows.Scan(&qrt.ID, &qrt.Title, &qrt.StartDate, &qrt.EndDate)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		quarters = append(quarters, qrt)
	}

	quartersJSON, err := json.Marshal(quarters)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Fprint(w, string(quartersJSON))
}
