/*
 * package: domain
 * --------------------
 * Includes definitions of types representing the domain model.
 * Also includes methods defined upon these types.
 */

package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strings"
)

type ToDoList struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description *string            `json:"description" bson:"description"`
	Tasks       []Task             `json:"tasks,omitempty" bson:"tasks,omitempty" validate:"required,dive,required"`
}

type Task struct {
	Id          string  `json:"id" bson:"id"`
	Name        string  `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description *string `json:"description" bson:"description"`
}

/*
 * Method: toDoList.AssignTaskIDs
 * --------------------
 * Assigns new, unique ids (uuid) to every Task in the ToDOList.
 * All existing Task ids will be overwritten.
 * Modifies the ToDoList it is applied to (pointer receiver)
 */

func (toDoList *ToDoList) AssignTaskIDs() {
	for i := range toDoList.Tasks {
		toDoList.Tasks[i].Id = uuid.NewString()
	}
}

/*
 * Method: toDoList.ResetID
 * --------------------
 * Resets id (primitive.ObjectID) to the zero value (ObjectID(000000000000000000000000)).
 * An existing id will be overwritten.
 * Modifies the ToDoList it is applied to (pointer receiver).
 *
 * returns: none
 */

func (toDoList *ToDoList) ResetID() {
	toDoList.Id = primitive.ObjectID{}
}

/*
 * Method: toDoList.Validate
 * --------------------
 * Validates the ToDoList using github.com/go-playground/validator/v10
 * Rules are defined in the tags provided in the ToDoList type definition.
 *
 * returns: a pointer to an errs.ValidationError with invalid fields
 *          and their respective violations in case of failed validation.
 *          Otherwise, on successful validation, nil is returned.
 */

func (toDoList ToDoList) Validate() *errs.ValidationError {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(toDoList)

	if err != nil {
		var invalidFields = make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := strings.SplitAfterN(e.Namespace(), ".", 2)[1]
			invalidFields[fieldName] = e.Tag()
		}
		return errs.NewValidationError(invalidFields)
	} else {
		return nil
	}
}
