/*
 * package: services
 * --------------------
 * Includes service implementation(s) (as defined in package ports)
 */

package services

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/core/ports"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
)

type DefaultToDoListService struct {
	repo ports.ToDoListRepository
}

/*
 * Method: DefaultToDoListService.GetAllLists
 * --------------------
 * Retrieves all ToDoLists using the injected repository and does not modify their order or applies filtering.
 *
 * returns: a pointer to a slice of domain.ToDoList and nil error in case of success.
 *          Otherwise nil and a pointer to an errs.AppError are returned.
 */

func (defaultToDoListService DefaultToDoListService) GetAllLists() (*[]domain.ToDoList, *errs.AppError) {
	lists, err := defaultToDoListService.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return lists, nil
}

/*
 * Method: DefaultToDoListService.SaveList
 * --------------------
 * Saves a list using the injected repository. The id is reset to its zero value
 * to overwrite a potentially existing client-side provided id and ensure proper
 * id assignment by the database.
 * Task ids are (re)assigned.
 *
 * newList: a domain.ToDoList intended for saving.
 *
 * returns: a pointer to a domain.ToDoList and nil error in case of success.
 *          Otherwise nil and a pointer to an errs.AppError are returned.
 */

func (defaultToDoListService DefaultToDoListService) SaveList(newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	newList.ResetID()
	newList.AssignTaskIDs()
	list, err := defaultToDoListService.repo.Save(newList)
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
 * Method: DefaultToDoListService.GetOneListById
 * --------------------
 * Retrieves a list with a provided id using the injected repository.
 *
 * id: a string representation of the requested list's object id
 *
 * returns: a pointer to a domain.ToDoList and nil error in case of success.
 *          Otherwise nil and a pointer to an errs.AppError are returned.
 */

func (defaultToDoListService DefaultToDoListService) GetOneListById(id string) (*domain.ToDoList, *errs.AppError) {
	list, err := defaultToDoListService.repo.GetOneById(id)
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
 * Method: DefaultToDoListService.UpdateOneListById
 * --------------------
 * Updates an existing list using the injected repository. The id is reset to its zero value
 * to overwrite a potentially existing client-side provided id and ensure proper
 * id assignment by the database.
 * Task ids are (re)assigned.
 *
 * id: a string representation of the object id belonging to the list intended to be updated.
 *
 * returns: a pointer to a domain.ToDoList and nil error in case of success.
 *          Otherwise nil and a pointer to an errs.AppError are returned.
 */

func (defaultToDoListService DefaultToDoListService) UpdateOneListById(id string, newList domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	newList.ResetID()
	newList.AssignTaskIDs()
	list, err := defaultToDoListService.repo.UpdateOneById(id, newList)
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
 * Method: DefaultToDoListService.DeleteListById
 * --------------------
 * Deletes an existing list using the injected repository.
 *
 * id: a string representation of the object id belonging to the list intended for deletion.
 *
 * returns: nil in case of success.
 *          Otherwise a pointer to an errs.AppError is returned.
 */

func (defaultToDoListService DefaultToDoListService) DeleteListById(id string) *errs.AppError {
	err := defaultToDoListService.repo.DeleteOneById(id)
	if err != nil {
		return err
	}
	return nil
}

/*
 * Function: NewToDoListService
 * --------------------
 * Instantiates a new DefaultToDoListService for dependency injection.
 *
 * repo: an implementation of ports.ToDoListRepository
 *
 * returns: an instance of DefaultToDoListService including the provided repository.
 */

func NewToDoListService(repo ports.ToDoListRepository) DefaultToDoListService {
	return DefaultToDoListService{repo}
}
