package database

import "github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"

type IUserDB interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type IProductDB interface {
	Create(product *entity.Product) error
	FindById(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
