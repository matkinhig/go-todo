package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllDataTodo(c echo.Context) error {
	fmt.Println("start process...")
	return c.JSON(http.StatusFound, "")
}
