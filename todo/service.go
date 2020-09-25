package todo

import (
	"errors"
)

type Service interface {
	GetAllForUser(username string) ([]Todo, error)
	GetByID(id string) (Todo, error)
	Create(req CreateRequest) (string, error)
	Update(id string, todo Todo) error
	Delete(id string) error
}

var (
	ErrInconsistentIDs = errors.New("Inconsistent IDs")
	ErrNotFound        = errors.New("Not found")
)

type service struct {
	repostory Repository
}

func NewService(rep Repository) Service {
	return &service{
		repostory: rep,
	}
}

func (s service) GetAllForUser(username string) (todos []Todo, err error) {

	todos, err = s.repostory.GetAllForUser(username)
	if todos == nil {
		todos = make([]Todo, 0)
	}
	return todos, err
}

func (s service) GetByID(id string) (Todo, error) {
	return s.repostory.GetByID(id)
}

func (s service) Create(req CreateRequest) (string, error) {
	return s.repostory.Create(req)

}

func (s service) Update(id string, todo Todo) error {
	return s.repostory.Update(id, todo)
}

func (s service) Delete(id string) error {
	return s.repostory.Delete(id)

}
