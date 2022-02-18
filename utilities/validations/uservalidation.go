package validations

import (
	"regexp"
	"radiant_cloud_assesment/models"
)

func ValidateUserData(user models.Users) (bool, string) {
	firstName, lastName, email := user.FirstName, user.LastName, user.Email
	if firstName == "" || lastName == "" {
		return false, "one or more field(s) is missing"
	}
	var reEmail = regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`)
	if !(reEmail.MatchString(email)) {
		return false, "Please provide a valid email"
	}
 
	return true, ""
}