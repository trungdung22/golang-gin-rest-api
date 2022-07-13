package services

import (
	"crud-api/models"
	"crud-api/validators"
)

func CreateOneUser(userDto *validators.UserSignUpRequest) error {
	var user models.UserModel
	user.Username = userDto.Username
	user.Bio = userDto.Bio
	user.Email = userDto.Email

	err := user.SetPassword(userDto.Password)

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

func FindOneUserByAttribute(condition interface{}) (models.UserModel, error) {
	handler := models.UserDaoHandler{}
	return handler.GetByCondition(condition)
}
