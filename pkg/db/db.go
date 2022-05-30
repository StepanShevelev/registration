package db

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	salt       = "hdssdvszxzad"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 60
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb() {
	dsn := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable timezone=Europe/Moscow"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		UppendErrorWithPath(err)
		//TODO Ты пытаешься записывать ошибку в базу до того, как подключился в базу.
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&ErrLogs{})

	Database = DbInstance{
		Db: db,
	}
}

func CreateUser(user *User) {

	Database.Db.Create(&user)

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
		//TODO не возвращаешь ошибки
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

func LoginCheck(email string, password string) error {

	var err error
	var u *User

	err = Database.Db.Model(&u).Where("email = ?", email).Take(&u).Error
	if err != nil {
		UppendErrorWithPath(err)
		//TODO не возвращаешь ошибки
	}

	newHash := GeneratePasswordHash(password)
	oldHash := u.PasswordHash
	if newHash != oldHash {
		log.Fatal("password is incorrect")
		//TODO не возвращаешь текста ошибки, то есть в случае происхождения ошибки вернул по сути ничего
		return nil
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

	result := Database.Db.Preload("Posts").Find(&user, "id = ?", Id)

	if result.Error != nil {
		//TODO не возвращаешь и не логируешь ошибку
		return nil, nil
	}
	return user, nil
}

func FindUserByEmail(Email string) (*User, error) {
	var user *User

	result := Database.Db.Find(&user, "email = ?", Email)

	if result.Error != nil {
		//TODO не возвращаешь и не логируешь ошибку
		return nil, nil
	}
	return user, nil
}

func CreatePost(post *Post) {

	Database.Db.Create(&post)
	//TODO не обрабатываешь возможные ошибки и не даёшь никакого ответа по результату
}
