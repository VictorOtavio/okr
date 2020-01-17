package main

import (
	"fmt"

	"github.com/VictorOtavio/4okr/internal/router"
	"github.com/VictorOtavio/4okr/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.Load()

	if viper.GetBool("app.debug") {
		gin.SetMode("debug")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	router := router.Routes()
	router.Run(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.host"),
		viper.GetString("server.port"),
	))
}
