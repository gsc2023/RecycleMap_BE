package domain

type User struct {
	Email         string
	EmailVerified bool
	PhoneNumber   string
	Password      string
	DisplayName   string
	PhotoURL      string
	Disabled      bool
}

type UserDto struct {
	ID   string
	User User
}
