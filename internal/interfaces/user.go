package interfaces

import (
	"github.com/IlyaZayats/shopus/internal/entity"
)

type UserRepository interface {
	Login(user entity.User) (int, error)
}
