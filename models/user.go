package models

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email;unique_index"`
	Bio      string `gorm:"column:bio;size:1024"`
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
	GetByCondition(condition interface{}) (UserModel, error)
	FetchUsers(pageSize int, page int) ([]UserModel, int, error)
}

type UserDaoHandler struct {
}

func (h UserDaoHandler) Create(data interface{}) error {
	fmt.Println(data)
	db := GetDB()
	err := db.Save(data).Error
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

func (h UserDaoHandler) GetUserByEmail(email string) (UserModel, error) {
	db := GetDB()
	var userModel UserModel
	err := db.Where("email = ?", email).First(&userModel).Error
	return userModel, err
}

func (h UserDaoHandler) GetByCondition(condition interface{}) (UserModel, error) {
	database := GetDB()
	var user UserModel
	err := database.Where(condition).First(&user).Error
	return user, err
}

func (h UserDaoHandler) FetchUsers(pageSize int, page int) ([]UserModel, int, error) {
	database := GetDB()
	var users []UserModel
	var count int64
	tx := database.Begin()
	database.Model(&users).Count(&count)

	tx.Model(&users).Offset((page - 1) * page).Limit(pageSize).Find(&users)
	err := tx.Commit().Error

	return users, int(count), err
}
