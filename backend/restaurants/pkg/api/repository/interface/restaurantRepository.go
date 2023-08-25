package _interface

type MenuRepository interface {
	AddMenuItemInMenu()
	RemoveMenuItemInMenu()
}
type MenuItemRepository interface {
	CreateMenuItem()
	GetMenuItemInMenu()
	GetMenuItemsInOrder()
	UpdateMenuItem()
	RemoveMenuItem()
	ChangeMenuItemAvailability()
}
type OrderRepository interface {
	CreateOrder()
	UpdateOrder()
	GetOrder()
}

type Repository struct {
	MenuRepository
	MenuItemRepository
	OrderRepository
}

func NewRepository(menuRepository MenuRepository, menuItemRepository MenuItemRepository, orderRepository OrderRepository) *Repository {
	return &Repository{MenuRepository: menuRepository, MenuItemRepository: menuItemRepository, OrderRepository: orderRepository}
}
