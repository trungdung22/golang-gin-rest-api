package services

import (
	"crud-api/models"
	"crud-api/utilities"
	"crud-api/validators"
)

func CreateOneUser(userDto *validators.UserSignUpRequest) (models.User, error) {
	user := models.User{
		Username: userDto.Username,
		Bio:      userDto.Bio,
		Email:    userDto.Email,
	}
	handler := models.UserDaoHandler{}
	if err := handler.Create(&user); err != nil {
		return user, err
	}
	return user, nil
}

func GetOrCreateUserByEmail(authPayload *utilities.GoogleUserPayload) (models.User, error) {
	user, err := GetUserByEmail(authPayload.Email)

	if err == nil {
		return user, nil
	}

	userDto := validators.UserSignUpRequest{
		Username: authPayload.Name,
		Email:    authPayload.Email,
		Bio:      authPayload.Locale,
	}
	return CreateOneUser(&userDto)
}

func GetUserByEmail(email string) (models.User, error) {
	handler := models.UserDaoHandler{}
	return handler.GetUserByEmail(email)
}

func GetUserById(id int) (models.User, error) {
	handler := models.UserDaoHandler{}
	return handler.GetById(id)
}

func FindOneUserByAttribute(condition interface{}) (models.User, error) {
	handler := models.UserDaoHandler{}
	return handler.GetByCondition(condition)
}

func FindUserPage(pageSize int, page int) ([]models.User, int, error) {
	handler := models.UserDaoHandler{}
	return handler.FetchUsers(pageSize, page)
}
