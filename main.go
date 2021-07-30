package main

import (
	"fmt"

	"farras/config"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("hello")
	e := echo.New()
	config.InitDB()
	e.Start(":8080")
}
