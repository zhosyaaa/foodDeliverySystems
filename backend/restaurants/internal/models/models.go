package models

type Dish struct {
	ID           uint64     `gorm:"primaryKey" json:"id"`
	RestaurantID uint64     `gorm:"column:restaurant_id" json:"restaurantId"`
	Name         string     `gorm:"column:name" json:"name"`
	Description  string     `gorm:"column:description" json:"description"`
	Availability uint64     `gorm:"column:availability" json:"availability"`
	Price        float64    `gorm:"column:price" json:"price"`
	Images       []string   `gorm:"column:images" json:"images"`
	Ingredients  []string   `gorm:"column:ingredients" json:"ingredients"`
	Categories   []Category `gorm:"many2many:dish_categories;" json:"categories"`
}

type Menu struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Dishes []Dish `gorm:"foreignKey:RestaurantID" json:"dishes"`
}

type Category struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}
type OrderItem struct {
	ID       uint   `gorm:"primaryKey"`
	DishID   uint64 `gorm:"column:dish_id"`
	Quantity int32  `gorm:"column:quantity"`
	OrderID  uint   `gorm:"index"`
}

type Location struct {
	ID         uint   `gorm:"primaryKey"`
	City       string `gorm:"column:city"`
	PostalCode string `gorm:"column:postal_code"`
	Address    string `gorm:"column:address"`
	Country    string `gorm:"column:country"`
}

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerID   string `gorm:"column:customer_id"`
	RestaurantID uint64 `gorm:"column:restaurant_id"`
	Items        []OrderItem
	TotalPrice   float64  `gorm:"column:total_price"`
	Status       string   `gorm:"column:status"`
	Location     Location `gorm:"foreignKey:LocationID"`
	LocationID   uint     `gorm:"column:location_id"`
}
