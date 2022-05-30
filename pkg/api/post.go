package api

import (
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type PostInput struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	UserId      int    `json:"user_id" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		mydb.UppendErrorWithPath(err)
		//TODO зачем продолжается функция после этой ошибки?
		//TODO нужен ответ сервера об ошибке. прим c.JSON(http.StatusBadRequest, gin.H{"error": "хочу жсон"})
	}
	var post mydb.Post
	user, err := mydb.FindUserById(input.UserId)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords doesn`t match "})
		return
	}

	post.Title = input.Title
	post.Description = input.Description
	post.Image = input.Image
	post.Users = append(post.Users, mydb.User{ID: input.UserId})

	mydb.Database.Db.Model(&post).Association("Users").Append(&user)
	mydb.CreatePost(&post)
}

func UpdatePost(ctx *gin.Context) {

	category, okCat := getCategoryById(id, w)
	if !okCat {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		return
	}

	mydb.Database.Db.Save(&category)

}

func DeletePost(ctx *gin.Context) {

	category, okCat := getCategoryById(id, w)
	if !okCat {
		return
	}

	mydb.Database.Db.Delete(&category)
	w.WriteHeader(http.StatusOK)
}
