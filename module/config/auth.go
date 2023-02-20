package config

import (
	"domain"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"firebase.google.com/go/auth/hash"
)

func B64URLdecode(s string) []byte {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		log.Println("URL: Failed to decode string", err)
	}

	return b
}

func B64Stddecode(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Println("String: Failed to decode string", err)
	}
	return b
}

func GetHashScrypt() *hash.Scrypt {
	data, err := os.Open("module/config/hashConfig.json")

	if err != nil {
		log.Println("Hash Scrypt", err)
	}

	byteValue, _ := ioutil.ReadAll(data)

	var stringScrypt domain.ScryptString

	json.Unmarshal(byteValue, &stringScrypt)

	hashScrypt := hash.Scrypt{
		Key:           B64Stddecode(stringScrypt.Key),
		SaltSeparator: B64Stddecode(stringScrypt.SaltSeparator),
		Rounds:        stringScrypt.Rounds,
		MemoryCost:    stringScrypt.MemoryCost,
	}

	return &hashScrypt
}
