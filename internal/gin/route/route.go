package route

import (
	"deploy-test/config"
	todo "deploy-test/internal/gin/handler/todo"
	"deploy-test/internal/gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(config *config.Config) *gin.Engine {
	if config.Gin.Mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.InitLogger)
	r.Use(middleware.Recovery)

	todoV1 := r.Group("/todos")
	{
		todoV1.POST("/add", todo.Add)
		todoV1.GET("/list", todo.List)
		todoV1.GET("/list/:id", todo.Get)
		todoV1.PATCH("/:id/edit", todo.Update)
		todoV1.PATCH("/:id/delete", todo.Delete)
	}

	return r
}
