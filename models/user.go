package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
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

type UserDao interface {
	Create(data interface{}) error
	Update(model *UserModel, data interface{}) error
	Delete(id int) error
	GetById(id int) (UserModel, error)
	SearchUser(username string, limit string, page string) ([]UserModel, error)
}

type Handler struct {
}

func (h Handler) Create(data interface{}) error {
	db := GetDB()
	err := db.Create(data).Error
	return err
}

func (h Handler) Update(model *UserModel, data interface{}) error {
	db := GetDB()
	err := db.Model(model).Updates(data).Error
	return err
}

func (h Handler) Delete(id int) error {
	db := GetDB()
	var userModel UserModel

	if err := db.Where("id = ?", id).First(&userModel).Error; err != nil {
		return err
	}

	err := db.Delete(&userModel).Error
	return err
}

func (h Handler) GetById(id int) (UserModel, error) {
	db := GetDB()
	var userModel UserModel
	err := db.Where("id = ?", id).First(&userModel).Error
	return userModel, err
}
