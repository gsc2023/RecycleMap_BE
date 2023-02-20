package domain

type SigninRequestDto struct {
	UID string
}

type ScryptString struct {
	Key           string
	SaltSeparator string
	Rounds        int
	MemoryCost    int
}
