package cabin

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/viper"
)

// Quarter representation in the OKR methodology
type Quarter struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// QuarterIndex lists all quarters
func QuarterIndex(ctx *gin.Context) {
	db, err := sql.Open("sqlite3", viper.GetString("database.sqlite.path"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows, err := db.Query("SELECT id, title, start_date, end_date FROM quarters")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	quarters := make([]Quarter, 0)

	for rows.Next() {
		var qrt Quarter
		err := rows.Scan(&qrt.ID, &qrt.Title, &qrt.StartDate, &qrt.EndDate)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		quarters = append(quarters, qrt)
	}

	ctx.JSON(200, quarters)
}
