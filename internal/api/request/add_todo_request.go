package request

import (
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func InitializeAddTodo(c echo.Context) (*command.AddTodo, error) {
	var addTodo command.AddTodo
	if err := c.Bind(&addTodo); err != nil {
		return &addTodo, errorx.Decorate(err, "failed to bind data")
	}

	if err := c.Validate(addTodo); err != nil {
		return &addTodo, err
	}
	return &addTodo, nil
}
