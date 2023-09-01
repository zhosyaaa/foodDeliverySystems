package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	Rating      float32    `gorm:"column:rating"`
	Locations   []Location `gorm:"many2many:restaurant_locations;"`
	Menu        Menu
}

type Location struct {
	gorm.Model
	City        string       `gorm:"column:city"`
	PostalCode  string       `gorm:"column:postal_code"`
	Address     string       `gorm:"column:address"`
	Country     string       `gorm:"column:country"`
	Restaurants []Restaurant `gorm:"many2many:restaurant_locations;"`
}

type Menu struct {
	gorm.Model
	Dishes []Dish `gorm:"foreignKey:RestaurantID"`
}

type Dish struct {
	gorm.Model
	RestaurantID uint64     `gorm:"column:restaurant_id"`
	Name         string     `gorm:"column:name"`
	Description  string     `gorm:"column:description"`
	Availability uint64     `gorm:"column:availability"`
	Price        float64    `gorm:"column:price"`
	Images       []string   `gorm:"column:images"`
	Ingredients  []string   `gorm:"column:ingredients"`
	Categories   []Category `gorm:"many2many:dish_categories;"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Dishes      []Dish `gorm:"many2many:dish_categories;"`
}
