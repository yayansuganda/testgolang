package models

import (
	"productservice/config"
)

type Product struct {
	NamaProduct string `json:"nama_product" form:"nama_product"`
	Harga       int    `json:"harga" form:"harga"`
	Stok        int    `json:"stok" form:"stok"`
}

func (product *Product) CreateProduct() error {
	if err := config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (product *Product) UpdateProduct(id string) error {
	if err := config.DB.Model(&Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

func (product *Product) DeleteProduct(id string) error {
	if err := config.DB.Where("id = ?", id).Delete(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductOneById(id string) (Product, error) {
	var products Product
	result := config.DB.Where("id = ?", id).First(&products)
	return products, result.Error
}

func GetProductAll() ([]Product, error) {
	var products []Product
	result := config.DB.Find(&products)

	return products, result.Error
}
