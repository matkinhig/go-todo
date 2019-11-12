package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetAllDataTodo(c echo.Context) error {
	fmt.Println("start process...")
	return c.JSON(http.StatusFound, "")
}
