package main

import (
	"server/math"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.GET("/add/:left/:right", math.AddHandler)
	e.Run(standard.New(":8000"))
}
