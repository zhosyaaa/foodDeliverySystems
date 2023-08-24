package models

import "gorm.io/gorm"

//message MenuItem {
//string id = 1;
//string name = 2;
//string description = 3;
//double price = 4;
//bool available = 5;
//}
//
//message Menu {
//repeated MenuItem items = 1;
//}
//
//message Order {
//string order_id = 1;
//repeated MenuItem items = 2;
//string status = 3;
//}
//
//service RestaurantService {
//rpc UpdateMenu(Menu) returns (Menu);
//rpc GetMenu(Empty) returns (Menu);
//rpc ProcessOrder(Order) returns (Order);
//}

type MenuItem struct {
	gorm.Model
	name        string  `json:"name"`
	description string  `json:"description"`
	price       float64 `json:"price"`
	available   int     `json:"available"`
}
type Menu struct {
	gorm.Model
	menuItems []MenuItem `json:"menuItems" gorm:"foreignKey:ID"`
}

// хотя можно разделить
type Order struct {
	gorm.Model
	menuItems []MenuItem `json:"menuItems" gorm:"foreignKey:ID"`
	status    string     `json:"status"`
}
