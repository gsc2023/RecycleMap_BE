package service

import (
	"domain"
	"log"
	"module/config"

	"firebase.google.com/go/auth"
)

func JoinUser(user domain.User) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL(user.PhotoURL).
		Disabled(user.Disabled)
	u, err := config.GetAuth().CreateUser(config.Ctx, params)
	if err != nil {
		log.Printf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	return u, err
}

func SignIn(request domain.SigninRequestDto) (string, error) {
	token, err := config.GetAuth().CustomToken(config.Ctx, request.UID)
	if err != nil {
		log.Printf("error minting custom token: %v\n", err)
	}

	log.Printf("Got custom token: %v\n", token)
	return token, err
}

func VerifyToken(accessToken domain.AccessTokenContainer) (*auth.Token, error) {
	token, err := config.GetAuth().VerifyIDToken(config.Ctx, accessToken.AccessToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	return token, err
}
