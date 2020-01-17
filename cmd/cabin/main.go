package main

import (
	"fmt"

	"github.com/VictorOtavio/4okr/internal/cabin"
	"github.com/VictorOtavio/4okr/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.Load()
}

func main() {
	router := gin.Default()

	router.GET("/quarters", cabin.QuarterIndex)
	router.GET("/objectives", cabin.ObjectivesIndex)
	router.GET("/key-results", cabin.KeyResultsIndex)

	router.Run(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.host"),
		viper.GetString("server.port"),
	))
}
