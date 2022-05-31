package main

import (
	"net/http"

	"github.com/StepanShevelev/registration/pkg/api"
	cfg "github.com/StepanShevelev/registration/pkg/config"
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {

	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		mydb.UppendErrorWithPath(err)
		//TODO Ты пытаешься записывать ошибку в базу до того, как подключился в базу.
	}

	mydb.ConnectToDb()
	//api.InitRoutes()

	r := gin.Default()

	auth := r.Group("/auth")

	auth.POST("/register", api.SignUp)
	auth.POST("/login", api.SignIn)

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

	http.ListenAndServe(":8000", nil)

}
