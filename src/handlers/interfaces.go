package handlers

import (
	"github.com/google/uuid"

	"growth-place/src/services/user"
)

// IUserService presents user services methods
type IUserService interface {
	Authorization(email, password string) (user.AuthorizationView, error)
	Create(login string, name, email *string, phone *string, password string) error
	PasswordEdit(id uuid.UUID, password string) error
}
