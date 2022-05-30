package api

import (
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

func SignUp(c *gin.Context) {

	var input signUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		mydb.UppendErrorWithPath(err)
	}
	var u mydb.User

	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords doesn`t match "})
		return
	}

	u.Name = input.Name
	u.PasswordHash = mydb.GeneratePasswordHash(input.Password)
	u.Password = input.Password
	u.Email = input.Email
	u.JwtToken, _ = mydb.GenerateToken(input.Email)

	mydb.CreateUser(&u)

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": u.JwtToken,
	})

}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		mydb.UppendErrorWithPath(err)
		return
	}

	err := mydb.LoginCheck(input.Email, input.Password)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		return
	}

	token, err := mydb.GenerateToken(input.Email)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func GetUserProfile(c *gin.Context) {

	usr := GetProfile(c)

	c.JSON(http.StatusOK, map[string]interface{}{
		"name":  usr.Name,
		"email": usr.Email,
		"posts": usr.Posts,
	})
}

func GetProfile(ctx *gin.Context) *mydb.User {
	userId, okId := parseId(ctx)
	if !okId {
		return nil
	}

	user, err := mydb.FindUserById(userId)
	if err != nil {
		mydb.UppendErrorWithPath(err)
	}
	return user
}
