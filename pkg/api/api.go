package api

import (
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

//func InitRoutes() *gin.Engine {
//	router := gin.Default()
//
//	auth := router.Group("/auth")
//	{
//		auth.POST("/sign-up", SignUp)
//		auth.POST("/sign-in", SignIn)
//	}
//
//	profile := router.Group("/profile")
//	{
//		{
//			profile.GET("/:id", GetUserProfile)
//		}
//	}
//
//	//post := router.Group("/post")
//	//{
//	//	post.POST("/", CreatePost)
//	//	post.PATCH("/:id", UpdatePost)
//	//	post.DELETE("/:id", DeletePost)
//	//}
//
//	return router
//}

func parseId(c *gin.Context) (int, bool) {
	keys, ok := c.Request.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {

		return 0, false
	}
	userId, err := strconv.Atoi(keys[0])
	if err != nil {

		mydb.UppendErrorWithPath(err)
	}
	return userId, true
}
