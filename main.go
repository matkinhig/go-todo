package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/matkinhig/go-todo/config"
)

func main() {
	fmt.Println("start golang")
	fmt.Println(config.Config)

	fmt.Println(db.Client)
	e := echo.New()
	route.All(e)
	err := e.Start(":9090")

	if err != nil {
		fmt.Println("Cannot start server", err)
	}
}
