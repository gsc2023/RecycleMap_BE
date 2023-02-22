package domain

type SigninRequestDto struct {
	UID string
}

type AccessTokenContainer struct {
	AccessToken string
}

type ScryptString struct {
	Key           string
	SaltSeparator string
	Rounds        int
	MemoryCost    int
}
