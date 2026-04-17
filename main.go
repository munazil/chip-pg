package chippg

import "log"

func GenAuthToken(email string, secret string) (string, error) {
	log.Println("GenAuthToken", email, secret)
	return "Testing 123", nil
}
