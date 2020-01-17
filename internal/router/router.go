package router

import (
	"github.com/VictorOtavio/4okr/internal/cabin"
	"github.com/gin-gonic/gin"
)

// Routes returns the app route list
func Routes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		quarters := v1.Group("/quarters")
		{
			quarters.GET("/", cabin.QuarterIndex)
			quarters.POST("/", cabin.QuarterStore)
			// quarters.GET("/:quarter", cabin.QuarterShow)
			// quarters.PATCH("/:quarter", cabin.QuarterUpdate)
			// quarters.DELETE("/:quarter", cabin.QuarterDelete)
		}

		objectives := v1.Group("/objectives")
		{
			objectives.GET("/", cabin.ObjectivesIndex)
			// objectives.POST("/", cabin.ObjectivesStore)
			// objectives.GET("/:objectives", cabin.ObjectivesShow)
			// objectives.PATCH("/:objectives", cabin.ObjectivesUpdate)
			// objectives.DELETE("/:objectives", cabin.ObjectivesDelete)
		}

		keyResults := v1.Group("/key-results")
		{
			keyResults.GET("/", cabin.KeyResultsIndex)
			// keyResults.POST("/", cabin.KeyResultsStore)
			// keyResults.GET("/:keyResults", cabin.KeyResultsShow)
			// keyResults.PATCH("/:keyResults", cabin.KeyResultsUpdate)
			// keyResults.DELETE("/:keyResults", cabin.KeyResultsDelete)
		}
	}

	return router
}
