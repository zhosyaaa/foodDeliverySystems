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
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
