package domain

import "errors"

type CustomerRepository interface {
	Store(customer Customer)
	FindByID(id int) Customer
}

type ItemRepository interface {
	Store(item Item)
	FindByID(id int) Item
}

type OrderRepository interface {
	Store(order Order)
	FindByID(id int) Order
}

type Customer struct {
	ID   int
	Name string
}

type Item struct {
	ID        int
	Name      string
	Value     float64
	Available bool
}

type Order struct {
	ID       int
	Customer Customer
	Items    []Item
}

func (order *Order) Add(item Item) error {
	return errors.New("not implemented yet")
}

func (order *Order) value() (sum float64) {
	return 0
}