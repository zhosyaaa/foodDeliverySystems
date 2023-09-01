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

type Category struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

type Menu struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Dishes []Dish `gorm:"foreignKey:RestaurantID" json:"dishes"`
}
type OrderItem struct {
	ID       uint   `gorm:"primaryKey"`
	DishID   uint64 `gorm:"column:dish_id"`
	Quantity int32  `gorm:"column:quantity"`
	OrderID  uint   `gorm:"index"`
}

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerID   string `gorm:"column:customer_id"`
	RestaurantID uint64 `gorm:"column:restaurant_id"`
	Items        []OrderItem
	TotalPrice   float64
	Status       string
}
