package api

import (
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type PostInput struct {
	ID          int    `json:"id"`
	Title       string `json:"title" `
	Description string `json:"description"`
	Image       string `json:"image"`
	//UserId      int             `json:"user_id" binding:"required"`
}

func CreatePost(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth header is really empty"})
		return
	}

	var user *mydb.User

	userHeader := mydb.Database.Db.Find(&user, "jwt_token = ?", header)
	if userHeader.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		mydb.UppendErrorWithPath(userHeader.Error)
		return
	}
	headerParts := strings.Split(header, ".")
	if len(headerParts) != 3 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		mydb.UppendErrorWithPath(userHeader.Error)
		c.Abort()
		return
	}

	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong with json"})
		return

	}
	var post mydb.Post

	post.Title = input.Title
	post.Description = input.Description
	post.Image = input.Image
	post.Users = append(post.Users, mydb.User{ID: user.ID})

	mydb.Database.Db.Model(&post).Association("Users").Append(&user)
	//TODO нет проверки на ошибку
	err := mydb.CreatePost(c, &post)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		c.AbortWithError(http.StatusUnauthorized, err)
	}
	c.JSON(http.StatusOK, gin.H{"success": "Post Created"})
}

func UpdatePost(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth header is really empty"})

		return
	}

	var user *mydb.User

	userHeader := mydb.Database.Db.Find(&user, "jwt_token = ?", header)
	if userHeader.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithError(http.StatusUnauthorized, userHeader.Error)
		mydb.UppendErrorWithPath(userHeader.Error)
		return
	}
	headerParts := strings.Split(header, ".")
	if len(headerParts) != 3 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		mydb.UppendErrorWithPath(userHeader.Error)
		c.AbortWithError(http.StatusUnauthorized, userHeader.Error)
		return
	}

	var input PostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "что-то не то с жсоном"})
		return
	}
	post, err := mydb.FindPostById(input.ID)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn`t find your post "})
		return
	}
	post.ID = input.ID
	post.Title = input.Title
	post.Description = input.Description
	post.Image = input.Image

	result := mydb.Database.Db.Save(&post)

	if result.Error != nil {
		mydb.UppendErrorWithPath(result.Error)
		c.AbortWithError(http.StatusUnauthorized, result.Error)
	}
	c.JSON(http.StatusOK, gin.H{"success": "Post Updated"})
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    post.ID,
		"title": post.Title,
		"desc":  post.Description,
		"image": post.Image,
	})

}

func DeletePost(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth header is really empty"})

		return
	}

	var user *mydb.User

	userHeader := mydb.Database.Db.Find(&user, "jwt_token = ?", header)
	if userHeader.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired, sing-in again"})
		mydb.UppendErrorWithPath(userHeader.Error)
		c.AbortWithError(http.StatusUnauthorized, userHeader.Error)
		return
	}
	headerParts := strings.Split(header, ".")
	if len(headerParts) != 3 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		mydb.UppendErrorWithPath(userHeader.Error)
		c.Abort()
		return
	}

	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "что-то не то с жсоном"})
		return
	}
	post, err := mydb.FindPostById(input.ID)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn`t find your post "})
		return
	}

	result := mydb.Database.Db.Delete(&post)
	if result.Error != nil {
		mydb.UppendErrorWithPath(result.Error)
		c.AbortWithError(http.StatusUnauthorized, result.Error)
	}
	c.JSON(http.StatusOK, gin.H{"success": "Post Deleted"})
}
