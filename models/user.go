package models

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username     string `gorm:"column:username"`
	Email        string `gorm:"column:email;unique_index"`
	Bio          string `gorm:"column:bio;size:1024"`
	PasswordHash string `gorm:"column:password;not null"`
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (user *UserModel) GenerateJwtToken() string {
	// jwt.New(jwt.GetSigningMethod("HS512"))
	jwt_token := jwt.New(jwt.SigningMethodHS512)

	jwt_token.Claims = jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}

type UserDao interface {
	Create(data interface{}) error
	Update(model *UserModel, data interface{}) error
	Delete(id int) error
	GetById(id int) (UserModel, error)
	SearchUser(username string, limit string, page string) ([]UserModel, error)
}

type UserDaoHandler struct {
}

func (h UserDaoHandler) Create(data interface{}) error {
	db := GetDB()
	err := db.Create(data).Error
	return err
}

func (h UserDaoHandler) Update(model *UserModel, data interface{}) error {
	db := GetDB()
	err := db.Model(model).Updates(data).Error
	return err
}

func (h UserDaoHandler) Delete(id int) error {
	db := GetDB()
	var userModel UserModel

	if err := db.Where("id = ?", id).First(&userModel).Error; err != nil {
		return err
	}

	err := db.Delete(&userModel).Error
	return err
}

func (h UserDaoHandler) GetById(id int) (UserModel, error) {
	db := GetDB()
	var userModel UserModel
	err := db.Where("id = ?", id).First(&userModel).Error
	return userModel, err
}
