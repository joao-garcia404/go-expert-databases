package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model            // add created_at, updated_at and deleted_at columns at products table
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Kitchen"}
	db.Create(&category)

	category2 := Category{Name: "Eletronics"}
	db.Create(&category2)

	// create product
	product := Product{
		Name:       "Cooking pot",
		Price:      90.00,
		Categories: []Category{category, category2},
	}

	db.Create(&product)

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category)
	}
}
