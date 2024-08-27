package util

import (
	"go-practics/internal/model"
)

// ContainsEqualCompt Returns true if the objects are equal, false otherwise.
func ContainsEqualCompt(input model.ComptItem, db []model.QkeDbComponent) bool {
	for _, v := range db {
		if v.ComponentName == input.Name && v.ComponentType == input.Type && v.ComponentVersion == input.Version {
			return true
		}
	}
	return false
}
