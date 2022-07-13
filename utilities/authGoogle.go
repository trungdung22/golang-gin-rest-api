package utilities

import (
	"context"
	"crud-api/config"

	"google.golang.org/api/idtoken"
)

type GoogleUserPayload struct {
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

func GetGoogleUserPayload(IdToken string) (*GoogleUserPayload, error) {
	config, _ := config.LoadConfig("..")
	payload, err := idtoken.Validate(context.Background(), IdToken, config.GoogleClientID)

	if err != nil {
		return nil, err
	}

	userBody := &GoogleUserPayload{
		Email:          payload.Claims["email"].(string),
		Verified_email: payload.Claims["email_verified"].(bool),
		Name:           payload.Claims["name"].(string),
		Given_name:     payload.Claims["given_name"].(string),
		Picture:        payload.Claims["picture"].(string),
		Locale:         payload.Claims["locale"].(string),
	}

	return userBody, nil
}
