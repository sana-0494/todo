package core

import (
	"context"
	"errors"
	"time"
)

type Todo struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
}

type TodoStore interface {
	Create(ctx context.Context, t Todo) (string, error)
	List(ctx context.Context) (Todo, error)
	GetById(ctx context.Context, id string) (Todo, error)
	Update(ctx context.Context, id string, td Todo) (Todo, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Todo, error)
}

type Service struct {
	store TodoStore
}

func NewService(s TodoStore) Service {
	return Service{
		store: s,
	}
}

func (s Service) Create(ctx context.Context, td Todo) string {
	createdtodo, _ := s.store.Create(ctx, td)
	return createdtodo
}

func (s Service) List(ctx context.Context) Todo {
	listTodo, _ := s.store.List(ctx)
	return listTodo
}

func (s Service) Update(ctx context.Context, id string, td Todo) Todo {
	//ADD CODE TO CHECK IF TODO EXISTS
	//GOING WITH SIMPLE IMPLIMENTATIONS NOW
	updatedTodo, _ := s.store.Update(ctx, id, td)
	return updatedTodo
}

func (s Service) Delete(ctx context.Context, id string) bool {
	err := s.store.Delete(ctx, id)
	return err == nil
	//Again no error handling
	//Just a simple db call. Assume there is no error in this righteous world
}

func (s Service) Restore(ctx context.Context, id string) (Todo, error) {
	_, err := s.store.GetById(ctx, id)
	if err != nil {
		return Todo{}, errors.New("could not find the requested taks")
	}
	restoredTodo, _ := s.store.Restore(ctx, id)
	return restoredTodo, nil
}
