package service

import "github.com/rajnarayan1729/goCrudPostgres/model"

type ProductService interface {
	CreateProduct(*model.Product) (int, error)
	GetProduct(*string) (*model.Product, error)
	GetAllProduct() ([]*model.Product, error)
	UpdatePrice(*model.Product) error
	DeleteProduct(*string) error
}
