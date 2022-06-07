package api

import (
	mydb "github.com/StepanShevelev/registration/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "email"
)

type signUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

// SignUp
//@Summary SignUp
// @Tags auth
// @Description create account
// @ID sing-up
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "account info"
// @Success 200 {object} map[string]interface{} "token"
// @Failure 400,401 {object} map[string]interface{}
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {

	var input signUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong with json"})
		mydb.UppendErrorWithPath(err)
		return
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

	err := mydb.CreateUser(c, &u)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		c.AbortWithError(http.StatusUnauthorized, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": u.JwtToken,
	})

}

type signInInput struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

// SignIn
//@Summary SignIn
// @Tags auth
// @Description sing-in
// @ID sing-in
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {object} map[string]interface{} "token"
// @Failure 400 {object} map[string]interface{}
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong with binding json"})
		mydb.UppendErrorWithPath(err)
		return
	}

	err := mydb.LoginCheck(c, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password"})
		mydb.UppendErrorWithPath(err)
		c.Abort()
		return
	}

	token, notOk := mydb.GenerateToken(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error, could not generate token": notOk,
		})
		mydb.UppendErrorWithPath(err)
		return
	}

	user, err := mydb.FindUserByEmail(input.Email)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn`t find user "})
		return
	}

	user.JwtToken = token
	mydb.Database.Db.Save(&user)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// GetUserProfile
//@Summary SignIn
// @Tags API
// @Description user profile
// @ID get-profile
// @Produce  json
// @Success 200 {object} mydb.User
// @Failure 400,401 {object} map[string]interface{}
// @Router /API/profile [get]
func GetUserProfile(c *gin.Context) {
	//var input findProfileInInput

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		return
	}

	var user *mydb.User

	userHeader := mydb.Database.Db.Find(&user, "jwt_token = ?", header)
	if userHeader.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		mydb.UppendErrorWithPath(userHeader.Error)
		return
	}

	usr, err := mydb.FindUserById(user.ID)
	if err != nil {
		mydb.UppendErrorWithPath(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn`t find user "})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    usr.ID,
		"name":  usr.Name,
		"email": usr.Email,
		"posts": usr.Posts,
	})
}

func UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
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

}
