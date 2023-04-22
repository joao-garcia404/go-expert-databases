package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model   // add created_at, updated_at and deleted_at columns at products table
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	// create product
	product := Product{
		Name:       "Macbook",
		Price:      12000.00,
		CategoryID: category.ID,
	}

	db.Create(&product)

	// create serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	})

	// Gettting products
	var products []Product

	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, product := range products {
		fmt.Println(product)
	}

	// Getting categories and their products
	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category)
	}
}
