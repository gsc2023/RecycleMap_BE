package service

import (
	"domain"
	"errors"
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
		return u, err
	}
	log.Printf("Successfully created user: %v\n", u)
	return u, err
}

func UpdateUser(token *auth.Token, user domain.User) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL(user.PhotoURL).
		Disabled(user.Disabled)
	u, err := config.GetAuth().UpdateUser(config.Ctx, token.UID, params)
	if err != nil {
		log.Printf("error updating user: %v\n", err)
		return u, err
	}
	log.Printf("Successfully updated user: %v\n", u)
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

func VerifyToken(accessToken domain.AccessTokenContainer) (*auth.Token, error) { // Test: 현재 UID 그대로 반환
	token := auth.Token{UID: accessToken.AccessToken}

	return &token, nil
}

func VerifyToken1(accessToken domain.AccessTokenContainer) (*auth.Token, error) {
	token, err := config.GetAuth().VerifyIDToken(config.Ctx, accessToken.AccessToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	return token, err
}

func IsOwner(status bool, err error) error {
	if err != nil {
		log.Printf("error owner: %v\n", err)
		return err
	}

	if !status {
		err = errors.New("user is not this owner")
		log.Printf("error owner: %v\n", err)
		return err
	}

	return nil
}
