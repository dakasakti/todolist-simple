package routes

import (
	"testing-code/handler/activity"
	"testing-code/handler/todo"

	"github.com/labstack/echo/v4"
)

func RouteActivity(e *echo.Echo, ha activity.HandlerActivity) {
	// middleware
	e.POST("activity-groups", ha.Create)
	e.GET("activity-groups", ha.Reads)
	e.PATCH("activity-groups/:id", ha.Edit)
	e.DELETE("activity-groups/:id", ha.Delete)
	e.GET("activity-groups/:id", ha.Read)
}

func RouteTodo(e *echo.Echo, ht todo.HandlerTodo) {
	// middleware
	e.POST("todo-items", ht.Create)
	e.GET("todo-items", ht.Reads)
	e.PATCH("todo-items/:id", ht.Edit)
	e.DELETE("todo-items/:id", ht.Delete)
	e.GET("todo-items/:id", ht.Read)
}
