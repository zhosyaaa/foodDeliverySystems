package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string   `gorm:"column:username" json:"username,omitempty"`
	Password string   `gorm:"column:password" json:"password,omitempty"`
	Email    string   `gorm:"column:email" json:"email,omitempty"`
	UserRole string   `gorm:"column:userRole" json:"userRole,omitempty"`
	Location Location `gorm:"foreignKey:LocationID" json:"location,omitempty"`
}

type Location struct {
	gorm.Model
	City       string `gorm:"column:city" json:"city,omitempty"`
	PostalCode string `gorm:"column:postalCode" json:"postalCode,omitempty"`
	Address    string `gorm:"column:address" json:"address,omitempty"`
	Country    string `gorm:"column:country" json:"country,omitempty"`
}
