package database

import (
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{
		DB: db,
	}
}

func (p *ProductDB) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductDB) FindById(id string) (*entity.Product, error) {
	var product entity.Product

	err := p.DB.Where("id = ?", id).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())

	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *ProductDB) Delete(id string) error {
	product, err := p.FindById(id)

	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
