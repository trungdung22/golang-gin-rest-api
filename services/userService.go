package services

import (
	"crud-api/models"
	"crud-api/validators"
)

func CreateOneUser(userDto *validators.UserSignUpRequest) error {
	user := models.UserModel{
		Username: userDto.Username,
		Email:    userDto.Email,
		Bio:      userDto.Bio,
	}
	err := user.setPassword(userDto.Password)

	if err != nil {
		return err
	}

	handler := models.UserDaoHandler{}
	if err := handler.Create(user); err != nil {
		return err
	}
	return nil
}

func GetUserById(id int) (models.UserModel, error) {
	handler := models.UserDaoHandler{}
	return handler.GetById(id)
}
