package activity

import (
	"fmt"
	"testing-code/entity"
	"testing-code/helper"
	"testing-code/repository/activity"

	"github.com/go-playground/validator/v10"
)

type serviceActivity struct {
	ra activity.RepositoryActivity
	v  *validator.Validate
}

func NewServiceActivity(ra activity.RepositoryActivity) *serviceActivity {
	return &serviceActivity{ra, validator.New()}
}

type ServiceActivity interface {
	Create(req entity.ActivityCreateRequest) (int, entity.Response)
	Read(id string) (int, entity.Response)
	Reads() (int, entity.Response)
	Edit(id string, req entity.ActivityUpdateRequest) (int, entity.Response)
	Delete(id string) (int, entity.Response)
}

func (sa *serviceActivity) Create(req entity.ActivityCreateRequest) (int, entity.Response) {
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

	// process
	data := entity.Activity{
		Email: req.Email,
		Title: req.Title,
	}

	// create
	res, err := sa.ra.Store(data)
	if err != nil {
		return 500, entity.Response{
			Status:  "failed",
			Message: "failed to create activity",
		}
	}

	// response
	return 201, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	}
}

func (sa *serviceActivity) Reads() (int, entity.Response) {
	res, err := sa.ra.Shows()
	if err != nil {
		return 404, entity.Response{
			Status:  "Not Found",
			Message: "failed to get data activity",
		}
	}

	return 200, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	}
}

func (sa *serviceActivity) Read(id string) (int, entity.Response) {
	req, err := helper.ParseParameter(id)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: "failed to get number parameter",
		}
	}

	res, err := sa.ra.Show(req)
	if err != nil {
		return 404, entity.Response{
			Status:  "Not Found",
			Message: "failed to get data activity",
		}
	}

	return 200, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	}
}

func (sa *serviceActivity) Edit(id string, req entity.ActivityUpdateRequest) (int, entity.Response) {
	reqId, err := helper.ParseParameter(id)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: "failed to get number parameter",
		}
	}

	data, err := sa.ra.Show(reqId)
	if err != nil {
		return 400, entity.Response{
			Status:  "failed",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
		}
	}

	data.Title = req.Title

	res, err := sa.ra.Update(data)
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

func (sa *serviceActivity) Delete(id string) (int, entity.Response) {
	req, err := helper.ParseParameter(id)
	if err != nil {
		return 400, entity.Response{
			Status:  "Bad Request",
			Message: "failed to get number parameter",
		}
	}

	err = sa.ra.Delete(req)
	if err != nil {
		return 404, entity.Response{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			Data:    entity.ObjectNil{},
		}
	}

	return 200, entity.Response{
		Status:  "Success",
		Message: "Success",
		Data:    entity.ObjectNil{},
	}
}
