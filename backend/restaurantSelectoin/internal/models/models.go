package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name        string   `gorm:"column:name" json:"name,omitempty"`
	Description string   `gorm:"column:description" json:"description,omitempty"`
	Rating      float32  `gorm:"column:rating" json:"rating,omitempty"`
	Location    Location `gorm:"foreignKey:RestaurantID" json:"location,omitempty"`
	Menu        Menu     `gorm:"foreignKey:RestaurantID" json:"menu,omitempty"`
}

type Location struct {
	gorm.Model
	RestaurantID uint64 `gorm:"column:restaurant_id" json:"-"`
	City         string `gorm:"column:city" json:"city,omitempty"`
	PostalCode   string `gorm:"column:postal_code" json:"postalCode,omitempty"`
	Address      string `gorm:"column:address" json:"address,omitempty"`
	Country      string `gorm:"column:country" json:"country,omitempty"`
}

type Menu struct {
	gorm.Model
	RestaurantID uint64 `gorm:"column:restaurant_id" json:"-"`
	Dishes       []Dish `gorm:"foreignKey:RestaurantID" json:"dishes,omitempty"`
}

type Dish struct {
	gorm.Model
	RestaurantID uint64     `gorm:"column:restaurant_id" json:"-"`
	Name         string     `gorm:"column:name" json:"name,omitempty"`
	Description  string     `gorm:"column:description" json:"description,omitempty"`
	Availability uint64     `gorm:"column:availability" json:"availability,omitempty"`
	Price        float64    `gorm:"column:price" json:"price,omitempty"`
	Images       []string   `gorm:"column:images" json:"images,omitempty"`
	Ingredients  []string   `gorm:"column:ingredients" json:"ingredients,omitempty"`
	Categories   []Category `gorm:"many2many:dish_categories;" json:"categories,omitempty"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
}
