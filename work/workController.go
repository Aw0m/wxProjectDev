package work

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Controller() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})
	v1 := e.Group("/work")
	{
		v1.POST("/createReport", func(context *gin.Context) {
			userID := context.GetHeader("userID")
			teamIdStr := context.PostForm("teamID")
			done := context.PostForm("done")
			toDo := context.PostForm("toDo")
			problem := context.PostForm("problem")
			ServiceCreateReport(userID, teamIdStr, done, toDo, problem, context)
		})

		v1.GET("/getReport", func(context *gin.Context) {
			repIdStr := context.Query("repID")
			ServiceGetReport(repIdStr, context)
		})
	}

	return e
}
