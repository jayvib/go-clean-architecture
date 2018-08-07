package usecases

import (
	"domain"
	"errors"
)

type UserRepository interface {
	Store(user User)
	FindByID(id int) User
}

type User struct {
	ID       int
	IsAdmin  bool
	Customer domain.Customer
}

type Item struct {
	ID    int
	Name  string
	Value float64
}

// Logger will be implemented in Infrastructure layer
type Logger interface {
	Log(args ...interface{})
}

type OrderInteractor struct {
	UserRepository UserRepository
	OrderRepository domain.OrderRepository
	ItemRepository domain.ItemRepository
	Logger Logger	// will not compose it since it is not meant to be a logger object
}

func (interactor *OrderInteractor) Items(userId, orderId int) ([]Item, error) {
	return nil, errors.New("not implemented yet")
}

func (interactor *OrderInteractor) Add(userId, orderId, itemId int) error {
	return errors.New("not implemented yet")
}

// AdminOrderInteractor has an order interactor method
type AdminOrderInteractor struct {
	OrderInteractor
}

func (interactor *AdminOrderInteractor) Add(userId, orderId, itemId int) error {
	return errors.New("not implemented yet")
}