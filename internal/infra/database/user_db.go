package database

import (
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		DB: db,
	}
}

func (u *UserDB) Create(user *entity.User) error {
	return nil
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	return nil, nil
}
