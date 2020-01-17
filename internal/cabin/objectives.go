package cabin

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/viper"
)

type Objective struct {
	ID          int    `json:"id"`
	QuarterID   int    `json:"quarter_id"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

func ObjectivesIndex(ctx *gin.Context) {
	db, _ := sql.Open("sqlite3", viper.GetString("database.sqlite.path"))
	rows, _ := db.Query("SELECT id, quarter_id, icon, description FROM objectives")

	objectives := make([]Objective, 0)

	for rows.Next() {
		var obj Objective
		err := rows.Scan(&obj.ID, &obj.QuarterID, &obj.Icon, &obj.Description)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		objectives = append(objectives, obj)
	}

	ctx.JSON(200, objectives)
}
