package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorOtavio/4okr/pkg/config"
)

type quarter struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func quarters(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	db, _ := sql.Open("sqlite3", config.GetEnv("DB_PATH", "../database.sqlite"))
	stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS quarters (id INTEGER PRIMARY KEY, description TEXT)")
	stmt.Exec()

	stmt, _ = db.Prepare("INSERT INTO quarters (description) VALUES (?)")
	stmt.Exec("Lorem ipsum dolor sit amet")

	rows, _ := db.Query("SELECT id, description FROM quarters")

	var quarters []quarter

	for rows.Next() {
		var q quarter
		err := rows.Scan(&q.ID, &q.Description)
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
