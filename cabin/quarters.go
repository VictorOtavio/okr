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

	stmt, _ := db.Prepare("INSERT INTO `quarters` (title, start_date, end_date) VALUES (?, ?, ?)")
	stmt.Exec("Lorem ipsum dolor sit amet", "2020-01-01 00:00:00", "2020-03-28 23:59:59")

	rows, _ := db.Query("SELECT id, title, start_date, end_date FROM quarters")

	var quarters []quarter

	for rows.Next() {
		var q quarter
		err := rows.Scan(&q.ID, &q.Title, &q.StartDate, &q.EndDate)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		quarters = append(quarters, q)
	}

	quartersJSON, err := json.Marshal(quarters)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Fprint(w, string(quartersJSON))
}
