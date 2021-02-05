package services

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/core/ports"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
)

type DefaultToDoListService struct {
	repo ports.ToDoListRepository
}

func (defaultToDoListService DefaultToDoListService) GetAllLists() (*[]domain.ToDoList, *errs.AppError) {
	lists, err := defaultToDoListService.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (defaultToDoListService DefaultToDoListService) SaveList(newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	newList.ResetID()
	newList.AssignTaskIDs()
	list, err := defaultToDoListService.repo.Save(newList)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (defaultToDoListService DefaultToDoListService) GetOneListById(id string) (*domain.ToDoList, *errs.AppError) {
	list, err := defaultToDoListService.repo.GetOneById(id)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (defaultToDoListService DefaultToDoListService) UpdateOneListById(id string, newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	newList.ResetID()
	newList.AssignTaskIDs()
	list, err := defaultToDoListService.repo.UpdateOneById(id, newList)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (defaultToDoListService DefaultToDoListService) DeleteListById(id string) *errs.AppError {
	err := defaultToDoListService.repo.DeleteOneById(id)
	if err != nil {
		return err
	}
	return nil
}

func NewToDoListService(repo ports.ToDoListRepository) DefaultToDoListService {
	return DefaultToDoListService{repo}
}
