package main

import (
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "github.com/StepanShevelev/registration/docs"
	"github.com/StepanShevelev/registration/pkg/api"
	cfg "github.com/StepanShevelev/registration/pkg/config"
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
)

// @title Registration
// @version 1.0
// @description Registration and interaction with post`s

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		logrus.Info(err)
	}

	mydb.ConnectToDb(config)
	//api.InitRoutes()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := r.Group("/auth")

	auth.POST("/sing-up", api.SignUp)
	auth.POST("/sing-in", api.SignIn)

	apii := r.Group("/API", api.UserIdentity)
	{
		apii.POST("/create_post", api.CreatePost)
		apii.PATCH("/update_post", api.UpdatePost)
		apii.DELETE("/delete_post", api.DeletePost)
		apii.GET("/profile", api.GetUserProfile)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + config.Port)

	http.ListenAndServe(":"+config.Port, nil)

}
