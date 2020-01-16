package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorOtavio/4okr/pkg/config"
)

type keyResult struct {
	ID          int    `json:"id"`
	ObjectiveID int    `json:"objective_id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func keyResults(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	db, _ := sql.Open("sqlite3", config.GetEnv("DB_PATH", "../database/database.sqlite"))
	rows, _ := db.Query("SELECT id, objective_id, description, done FROM key_results")

	keyResults := make([]keyResult, 0)

	for rows.Next() {
		var kr keyResult
		err := rows.Scan(&kr.ID, &kr.ObjectiveID, &kr.Description, &kr.Done)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		keyResults = append(keyResults, kr)
	}

	keyResultsJSON, err := json.Marshal(keyResults)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Fprint(w, string(keyResultsJSON))
}
