package phonenumber

import "strconv"

func IsValid(phoneNumber string) bool {
	// TODO - we can use regular expression to support +98
	if len(phoneNumber) != 11 {

		return false
	}

	if phoneNumber[:2] != "09" {

		return false
	}

	if _, cErr := strconv.Atoi(phoneNumber); cErr != nil {

		return false
	}

	return true
}
