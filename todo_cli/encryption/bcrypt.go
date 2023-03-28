package encryption

import "golang.org/x/crypto/bcrypt"

func HashThePassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
