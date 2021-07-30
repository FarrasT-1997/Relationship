package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string
	Description string
	Books       []Book
}

type Book struct {
	gorm.Model
	Title       string
	CategoryID  uint
	Description Book_description
}

type Book_description struct {
	gorm.Model
	Description string
	BookID      uint
}

type Transaction struct {
	gorm.Model
	Status         string
	Payment_method string
	Cart_shops     Cart_shop
}

type Cart_shop struct {
	gorm.Model
	TransactionID uint
	BookID        uint
}
