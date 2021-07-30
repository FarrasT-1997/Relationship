package config

import (
	"farras/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	connectionString := "root:12345@tcp(localhost:3306)/alta_api_percobaan?parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Category{}, &model.Book{}, &model.Book_description{}, &model.Transaction{}, &model.Cart_shop{})

	//cari category fiksi one to many

	var category model.Category
	if err := db.Where("name=?", "non fiksi").First(&category).Error; err != nil {
		panic(err)
	}
	//cari book yang category_id nya sesuai dengan category
	var books []model.Book
	if err := db.Find(&books, "category_id=?", category.ID).Error; err != nil {
		panic(err)
	}

	//show description every one to one
	var books_2 model.Book
	if err := db.Where("id=?", 5).First(&books_2).Error; err != nil {
		panic(err)
	}

	var bookDescription model.Book_description
	if err := db.Where("book_id=?", books_2.ID).First(&bookDescription).Error; err != nil {
		panic(err)
	}

	//PRINT ALL
	fmt.Printf("CATEGORY : %s\n", category.Name)
	fmt.Printf("Deskripsi %s : %s\n", category.Name, category.Description)

	for index, book := range books {
		fmt.Printf("Book %d: %s\n", index+1, book.Title)
	}
	fmt.Println()
	fmt.Printf("Book's Title: %s\n", books_2.Title)
	fmt.Printf("Book's Category: %s\n", category.Name)
	fmt.Printf("Description : %s\n", bookDescription.Description)
}
