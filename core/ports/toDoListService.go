/*
 * package: ports
 * --------------------
 * Includes interface definitions for layer connections (services and repositories).
 */

package ports

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
)

//go:generate mockgen -destination=../../mocks/ports/mockToDoListService.go -package=ports github.com/luschnat-ziegler/toDoListAPI/core/ports ToDoListService
type ToDoListService interface {
	GetAllLists() (*[]domain.ToDoList, *errs.AppError)
	SaveList(domain.ToDoList) (*domain.ToDoList, *errs.AppError)
	GetOneListById(string) (*domain.ToDoList, *errs.AppError)
	UpdateOneListById(string, domain.ToDoList) (*domain.ToDoList, *errs.AppError)
	DeleteListById(string) *errs.AppError
}
