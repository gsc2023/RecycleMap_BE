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

func SignIn(request domain.SigninRequestDto) string {
	token, err := config.GetAuth().CustomToken(config.Ctx, request.UID)
	if err != nil {
		log.Printf("error minting custom token: %v\n", err)
	}

	log.Printf("Got custom token: %v\n", token)
	return token
}
