package database

import (
	"assignment-2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	var host = "localhost"
	var user = "postgres"
	var password = "1234"
	var dbPort = "5432"
	var dbName = "assignment_2"

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to database :", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(model.Order{}, model.Item{})

	if err != nil {
		fmt.Println("error on migration :", err)
		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}

func (d Database) CreateOrder(order model.Order) (model.Order, error) {
	newOrder := order

	err := d.db.Create(&newOrder).Error

	return newOrder, err
}

func (d Database) GetOrders() ([]model.Order, error) {
	var orders []model.Order

	err := d.db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting orders: ", err.Error())
		return nil, err
	}

	return orders, nil
}

func (d Database) UpdateOrder(id uint, order model.Order) (model.Order, error) {
	newOrder := order
	var orderTemplate model.Order

	for i := range newOrder.Items {
		result := d.db.Model(&newOrder.Items[i]).Where("item_id = ?", newOrder.Items[i].Item_ID).Updates(&newOrder.Items[i])
		if result.Error != nil {
			fmt.Println("error not find items: ", result.Error)
			return model.Order{}, result.Error
		} else if result.RowsAffected < 1 {
			return model.Order{}, fmt.Errorf("failed to update items. item_id=%d not found", newOrder.Items[i].Item_ID)
		}
	} 
	
	orderTemplate.CustomerName = newOrder.CustomerName
	orderTemplate.OrderedAt = newOrder.OrderedAt
	
	result := d.db.Model(&orderTemplate).Where("order_id=?", id).Updates(&orderTemplate)
	if result.Error != nil {
		fmt.Println("error not find model: ", result.Error)
		return model.Order{}, result.Error
	} else if result.RowsAffected < 1 {
		return model.Order{}, fmt.Errorf("failed to update order. order_id=%d not found", id)
	}

	err := d.db.Preload("Items").Where("order_id=?", id).Find(&newOrder).Error
	if err != nil {
		fmt.Println("error not find order with items: ", err)
		return model.Order{}, err
	}

	return newOrder, nil
}

func (d Database) DeleteOrder(id uint) (error) {
	if result := d.db.Where("order_id = ?", id).Delete(&model.Item{}); result.Error != nil {
		fmt.Println("Error deleting items :", result.Error)
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("failed to delete data. order_id=%d not found", id)
	}
	
	if result := d.db.Delete(&model.Order{}, id); result.Error != nil {
		fmt.Println("error deleting order :", result.Error)
		return result.Error
		} else if result.RowsAffected < 1 {
			return fmt.Errorf("failed to delete data. order_id=%d not found", id)
	}
	
	return nil
}