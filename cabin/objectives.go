package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorOtavio/4okr/pkg/config"
)

type objective struct {
	ID          int    `json:"id"`
	QuarterID   int    `json:"quarter_id"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

func objectives(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (*req).Method == "OPTIONS" {
		return
	}

	db, _ := sql.Open("sqlite3", config.GetEnv("DB_PATH", "../database/database.sqlite"))
	rows, _ := db.Query("SELECT id, quarter_id, icon, description FROM objectives")

	objectives := make([]objective, 0)

	for rows.Next() {
		var obj objective
		err := rows.Scan(&obj.ID, &obj.QuarterID, &obj.Icon, &obj.Description)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		objectives = append(objectives, obj)
	}

	objectivesJSON, err := json.Marshal(objectives)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Fprint(w, string(objectivesJSON))
}
