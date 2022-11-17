package todo

import (
	"testing-code/entity"
	"testing-code/services/todo"

	"github.com/labstack/echo/v4"
)

type handlerTodo struct {
	st todo.ServiceTodo
}

func NewHandlerTodo(st todo.ServiceTodo) *handlerTodo {
	return &handlerTodo{st}
}

type HandlerTodo interface {
	Create(c echo.Context) error
	Reads(c echo.Context) error
	Read(c echo.Context) error
	Edit(c echo.Context) error
	Delete(c echo.Context) error
}

func (ht *handlerTodo) Create(c echo.Context) error {
	var req entity.TodoCreateRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  "Failed",
			Message: "Failed binding data",
		})
	}

	code, res := ht.st.Create(req)
	return c.JSON(code, res)
}

func (ht *handlerTodo) Reads(c echo.Context) error {
	code, res := ht.st.Reads()
	return c.JSON(code, res)
}

func (ht *handlerTodo) Read(c echo.Context) error {
	id := c.Param("id")

	code, res := ht.st.Read(id)
	return c.JSON(code, res)
}

func (ht *handlerTodo) Edit(c echo.Context) error {
	id := c.Param("id")
	var req entity.TodoUpdateRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  "failed",
			Message: "failed binding data",
		})
	}

	code, res := ht.st.Edit(id, req)
	return c.JSON(code, res)
}

func (ht *handlerTodo) Delete(c echo.Context) error {
	id := c.Param("id")

	code, res := ht.st.Delete(id)
	return c.JSON(code, res)
}
