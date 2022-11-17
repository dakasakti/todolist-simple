package main

import (
	"fmt"
	"testing-code/config"
	"testing-code/entity"
	"testing-code/handler/activity"
	"testing-code/handler/todo"
	ra "testing-code/repository/activity"
	sa "testing-code/services/activity"

	rt "testing-code/repository/todo"
	st "testing-code/services/todo"

	"testing-code/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		status := `<p>Status OK.</p>`
		text := `<p>Created by. <a href="https://github.com/dakasakti">Mahmuda Karima</a></p>`
		return c.HTML(200, status+text)
	})

	// config db
	conf := config.GetConfig()
	db := config.InitMySQL(conf)
	db.AutoMigrate(&entity.Activity{}, &entity.Todo{})

	// activity
	ra := ra.NewRepositoryActivity(db)
	sa := sa.NewServiceActivity(ra)
	ha := activity.NewHandlerActivity(sa)

	// todo
	rt := rt.NewRepositoryTodo(db)
	st := st.NewServiceTodo(rt, ra)
	ht := todo.NewHandlerTodo(st)

	routes.RouteActivity(e, ha)
	routes.RouteTodo(e, ht)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.GetConfig().Port)))
}
