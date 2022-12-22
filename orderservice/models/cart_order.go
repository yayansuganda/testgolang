package models

import (
	"orderservice/config"
)

type Cart_order struct {
	ProductId int    `json:"product_id" form:"product_id"`
	Jumlah    string `json:"jumlah" form:"jumlah"`
}

func (cart_order *Cart_order) CreateCartOrder() error {
	if err := config.DB.Create(cart_order).Error; err != nil {
		return err
	}
	return nil
}

func (cart_order *Cart_order) UpdateCartOrder(id string) error {
	if err := config.DB.Model(&Cart_order{}).Where("id = ?", id).Updates(cart_order).Error; err != nil {
		return err
	}
	return nil
}

func (cart_order *Cart_order) DeleteCartOrder(id string) error {
	if err := config.DB.Where("id = ?", id).Delete(cart_order).Error; err != nil {
		return err
	}
	return nil
}

func GetOneById(id string) (Cart_order, error) {
	var cart_order Cart_order
	result := config.DB.Where("id = ?", id).First(&cart_order)
	return cart_order, result.Error
}

func GetAll() ([]Cart_order, error) {
	var cart_orders []Cart_order
	result := config.DB.Find(&cart_orders)

	return cart_orders, result.Error
}
