package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // add created_at, updated_at and deleted_at columns at products table
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.00,
	})

	// create batch
	products := []Product{
		{Name: "Macbook", Price: 12000.00},
		{Name: "Mouse", Price: 120.00},
		{Name: "Keychron", Price: 800.00},
	}

	db.Create(&products)

	// select one
	var product Product
	db.First(&product, 3)
	db.First(&product, "name = ?", "Mouse")

	// Select all
	var products1 []Product

	db.Find(&products1)

	for _, product := range products1 {
		fmt.Println(product)
	}

	// Select with limit and offset
	var products2 []Product

	db.Limit(2).Offset(2).Find(&products2)

	for _, product := range products {
		fmt.Println(product)
	}

	// Where
	var products3 []Product

	db.Where("price > ?", 1000).Find(&products3)

	fmt.Println(products3)

	// Like
	var products4 []Product

	db.Where("name LIKE ?", "%k%").Find(&products4)

	fmt.Println(products4)

	// insert and soft_delete
	var p Product
	db.First(&p, 1)
	p.Name = "New Macbook"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	db.Delete(p2)
}
