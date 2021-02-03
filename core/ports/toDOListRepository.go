package ports

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
)

type ToDoListRepository interface {
	GetAll() (*[]domain.ToDoList, *errs.AppError)
	GetOneById(string) (*domain.ToDoList, *errs.AppError)
	UpdateOneById(string, domain.ToDoList) (*domain.ToDoList, *errs.AppError)
	Save(domain.ToDoList) (*domain.ToDoList, *errs.AppError)
	DeleteOneById(string) (*int64, *errs.AppError)
}
