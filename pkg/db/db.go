package db

import (
	"crypto/sha1"
	"fmt"
	cfg "github.com/StepanShevelev/registration/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

const (
	salt       = "hdssdvszxzad"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 1
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb(config *cfg.Config) {
	//dsn := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable timezone=Europe/Moscow"

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=Europe/Moscow",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Info(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&ErrLogs{})

	Database = DbInstance{
		Db: db,
	}
}

func CreateUser(c *gin.Context, user *User) error {

	result := Database.Db.Create(&user)
	if result.Error != nil {
		UppendErrorWithPath(result.Error)
		c.AbortWithError(http.StatusUnauthorized, result.Error)
	}
	c.JSON(http.StatusOK, gin.H{"success": "user created"})
	return nil

}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateToken(email string) (string, error) {
	user, err := FindUserByEmail(email)
	if err != nil {
		UppendErrorWithPath(err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenTTL*24) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Email,
	})

	return token.SignedString([]byte(signingKey))
}

func LoginCheck(c *gin.Context, email string, password string) error {

	var u *User

	err := Database.Db.Model(&u).Where("email = ?", email).Take(&u).Error
	if err != nil {
		UppendErrorWithPath(err)
		return err
	}

	oldHash := u.PasswordHash
	newHash := GeneratePasswordHash(password)

	if newHash != oldHash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		c.AbortWithError(http.StatusUnauthorized, err)
		return err
	}

	//token, err := GenerateToken(u.Email)
	//
	//if err != nil {
	//	return "", err
	//}

	return err

}

func FindUserById(Id int) (*User, error) {
	var user *User

	result := Database.Db.Preload("Posts.Users").Find(&user, "id = ?", Id)

	if result.Error != nil {
		UppendErrorWithPath(result.Error)
		return nil, result.Error
	}
	return user, nil
}

func FindPostById(Id int) (*Post, error) {
	var post *Post

	result := Database.Db.Preload("Users").Find(&post, "id = ?", Id)

	if result.Error != nil {
		UppendErrorWithPath(result.Error)
		return nil, result.Error
	}
	return post, nil
}

func FindUserByEmail(Email string) (*User, error) {
	var user *User

	result := Database.Db.Find(&user, "email = ?", Email)

	if result.Error != nil {
		UppendErrorWithPath(result.Error)
		return nil, result.Error
	}
	return user, nil
}

func CreatePost(c *gin.Context, post *Post) error {

	result := Database.Db.Create(&post)
	if result.Error != nil {
		UppendErrorWithPath(result.Error)
		c.AbortWithError(http.StatusUnauthorized, result.Error)
	}
	return nil
}
