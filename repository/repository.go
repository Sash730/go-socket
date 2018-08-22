package repository

import (
	"github.com/Sash730/go-socket/model"
)

type PreferencesRepository interface {
	GetByUser(user model.User) (*model.UserPreferences, error)
}
