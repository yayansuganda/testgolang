package models

import (
	"orderservice/config"
)

type Order struct {
	ProductId int    `json:"product_id" form:"product_id"`
	Jumlah    int    `json:"jumlah" form:"jumlah"`
	Harga     string `json:"harga" form:"harga"`
}

func (order *Order) CreateOrder() error {
	if err := config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (order *Order) UpdateOrder(id string) error {
	if err := config.DB.Model(&Order{}).Where("id = ?", id).Updates(order).Error; err != nil {
		return err
	}
	return nil
}

func (product *Product) UpdateStok(id int, stok int) error {

	config.DB.Model(&Product{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"stok": stok,
		},
	)
	// if err := config.DB.Table("products").Where("id = ?", id).Updates(product).Error; err != nil {
	// 	return err
	// }
	return nil
}

func (order *Order) DeleteOrder(id string) error {
	if err := config.DB.Where("id = ?", id).Delete(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderOneById(id string) (Order, error) {
	var order Order
	result := config.DB.Where("id = ?", id).First(&order)
	return order, result.Error
}

func GetOrderAll() ([]Order, error) {
	var orders []Order
	result := config.DB.Table("orders").Find(&orders)

	return orders, result.Error
}

func GetProduct(id int) (Product, error) {
	var products Product
	result := config.DB.Where("id = ?", id).Select("id").First(&products)
	return products, result.Error
}

type Product struct {
	NamaProduct int `json:"nama_product" form:"nama_product"`
	Harga       int `json:"harga" form:"harga"`
	Stok        int `json:"stok" form:"stok"`
}
