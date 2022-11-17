package activity

import (
	"testing-code/entity"
	"testing-code/services/activity"

	"github.com/labstack/echo/v4"
)

type handlerActivity struct {
	sa activity.ServiceActivity
}

func NewHandlerActivity(sa activity.ServiceActivity) *handlerActivity {
	return &handlerActivity{sa}
}

type HandlerActivity interface {
	Create(c echo.Context) error
	Reads(c echo.Context) error
	Read(c echo.Context) error
	Edit(c echo.Context) error
	Delete(c echo.Context) error
}

func (ha *handlerActivity) Create(c echo.Context) error {
	var req entity.ActivityCreateRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  "Failed",
			Message: "Failed binding data",
		})
	}

	code, res := ha.sa.Create(req)
	return c.JSON(code, res)
}

func (ha *handlerActivity) Reads(c echo.Context) error {
	code, res := ha.sa.Reads()
	return c.JSON(code, res)
}

func (ha *handlerActivity) Read(c echo.Context) error {
	id := c.Param("id")

	code, res := ha.sa.Read(id)
	return c.JSON(code, res)
}

func (ha *handlerActivity) Edit(c echo.Context) error {
	id := c.Param("id")
	var req entity.ActivityUpdateRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  "failed",
			Message: "failed binding data",
		})
	}

	code, res := ha.sa.Edit(id, req)
	return c.JSON(code, res)
}

func (ha *handlerActivity) Delete(c echo.Context) error {
	id := c.Param("id")

	code, res := ha.sa.Delete(id)
	return c.JSON(code, res)
}
