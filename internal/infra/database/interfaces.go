package database

import "github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"

type UserDBInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
