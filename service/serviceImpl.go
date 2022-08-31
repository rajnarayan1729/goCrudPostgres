package service

import (
	"github.com/rajnarayan1729/goCrudPostgres/model"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB *gorm.DB
}

func NewServiceImpl(db *gorm.DB) ProductService {

	return &ProductServiceImpl{
		DB: db,
	}
}

func (si *ProductServiceImpl) CreateProduct(product *model.Product) (int, error) {
	result := si.DB.Create(product)

	if result.Error != nil {
		return -1, result.Error
	}

	return int(product.ID), nil

}

func (si *ProductServiceImpl) GetProduct(code *string) (*model.Product, error) {

	var product *model.Product
	result := si.DB.Where("code = ?", code).First(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (si *ProductServiceImpl) GetAllProduct() ([]*model.Product, error) {

	var products []*model.Product
	result := si.DB.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil

}

func (si *ProductServiceImpl) UpdatePrice(updateProduct *model.Product) error {

	var product *model.Product
	result := si.DB.Model(&product).Where("code = ?", updateProduct.Code).Updates(model.Product{Code: updateProduct.Code, Price: updateProduct.Price})

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (si *ProductServiceImpl) DeleteProduct(code *string) error {

	var product *model.Product
	result := si.DB.Where("code = ?", code).Delete(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
