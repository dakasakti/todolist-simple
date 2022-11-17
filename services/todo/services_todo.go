package todo

import (
	"fmt"
	"testing-code/entity"
	"testing-code/helper"
	"testing-code/repository/activity"
	"testing-code/repository/todo"

	"github.com/go-playground/validator/v10"
)

type serviceTodo struct {
	rt todo.RepositoryTodo
	ra activity.RepositoryActivity
	v  *validator.Validate
}

func NewServiceTodo(rt todo.RepositoryTodo, ra activity.RepositoryActivity) *serviceTodo {
	return &serviceTodo{rt, ra, validator.New()}
}

type ServiceTodo interface {
	Create(req entity.TodoCreateRequest) (int, entity.Response)
	Read(id string) (int, entity.Response)
	Reads() (int, entity.Response)
	Edit(id string, req entity.TodoUpdateRequest) (int, entity.Response)
	Delete(id string) (int, entity.Response)
}

func (sa *serviceTodo) Create(req entity.TodoCreateRequest) (int, entity.Response) {
	// validate
	err := sa.v.Struct(req)
	if err != nil {
		error := helper.ParseMessageErrors(err.(validator.ValidationErrors))

		return 400, entity.Response{
			Status:  "Bad Request",
			Message: error,
			Data:    entity.ObjectNil{},
		}
	}

	// check activity
	_, err = sa.ra.Show(req.ActivityGroupID)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: fmt.Sprintf("Todo with ID %d Not Found", req.ActivityGroupID),
		}
	}

	data := entity.Todo{
		Title:           req.Title,
		ActivityGroupID: req.ActivityGroupID,
	}

	// create
	res, err := sa.rt.Store(data)
	if err != nil {
		return 500, entity.Response{
			Status:  "failed",
			Message: "failed to create Todo",
		}
	}

	// response
	return 201, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	}
}

func (sa *serviceTodo) Reads() (int, entity.Response) {
	res, err := sa.rt.Shows()
	if err != nil {
		return 404, entity.Response{
			Status:  "Not Found",
			Message: "failed to get data Todo",
		}
	}

	return 200, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	}
}

func (sa *serviceTodo) Read(id string) (int, entity.Response) {
	req, err := helper.ParseParameter(id)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: "failed to get number partmeter",
		}
	}

	res, err := sa.rt.Show(req)
	if err != nil {
		return 404, entity.Response{
			Status:  "Not Found",
			Message: "failed to get data Todo",
		}
	}

	return 200, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	}
}

func (sa *serviceTodo) Edit(id string, req entity.TodoUpdateRequest) (int, entity.Response) {
	reqId, err := helper.ParseParameter(id)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: "failed to get number partmeter",
		}
	}

	data, err := sa.rt.Show(reqId)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: fmt.Sprintf("Todo with ID %s Not Found", id),
		}
	}

	data.Title = req.Title
	data.IsActive = req.IsActive

	res, err := sa.rt.Update(data)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: "failed to update asset",
		}
	}

	return 200, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	}
}

func (sa *serviceTodo) Delete(id string) (int, entity.Response) {
	req, err := helper.ParseParameter(id)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: "failed to get number partmeter",
		}
	}

	err = sa.rt.Delete(req)
	if err != nil {
		return 404, entity.Response{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %s Not Found", id),
			Data:    entity.ObjectNil{},
		}
	}

	return 200, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    entity.ObjectNil{},
	}
}
