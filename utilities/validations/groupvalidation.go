package validations

import (
	"radiant_cloud_assesment/models"
)

func ValidateGroupData(group models.Groups) (bool, string) {
	groupName := group.GroupName
	if groupName == "" {
		return false, "group name is required"
	}
	
	return true, ""
}