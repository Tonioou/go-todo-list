package request

import (
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func InitializeUpdateTodo(c echo.Context) (*command.UpdateTodo, error) {
	var updateTodo command.UpdateTodo
	if err := c.Bind(&updateTodo); err != nil {
		return &updateTodo, errorx.Decorate(err, "failed to bind data")
	}

	if err := c.Validate(updateTodo); err != nil {
		return &updateTodo, err
	}
	return &updateTodo, nil
}
