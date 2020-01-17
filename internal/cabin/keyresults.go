package cabin

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/viper"
)

type KeyResult struct {
	ID          int    `json:"id"`
	ObjectiveID int    `json:"objective_id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func KeyResultsIndex(ctx *gin.Context) {
	db, _ := sql.Open("sqlite3", viper.GetString("database.sqlite.path"))
	rows, _ := db.Query("SELECT id, objective_id, description, done FROM key_results")

	keyResults := make([]KeyResult, 0)

	for rows.Next() {
		var kr KeyResult
		err := rows.Scan(&kr.ID, &kr.ObjectiveID, &kr.Description, &kr.Done)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		keyResults = append(keyResults, kr)
	}

	ctx.JSON(200, keyResults)
}
