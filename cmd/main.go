package main

import (
	"github.com/StepanShevelev/registration/pkg/api"
	cfg "github.com/StepanShevelev/registration/pkg/config"
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		mydb.UppendErrorWithPath(err)
	}

	mydb.ConnectToDb()
	api.InitRoutes()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + config.Port)

	http.ListenAndServe(":8000", nil)

}
