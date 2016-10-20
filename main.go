package main

import (
	"github.com/ksmithut/go-math-server/math"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.GET("/add/:left/:right", math.AddHandler)
	e.GET("/subtract/:left/:right", math.SubtractHandler)
	e.Run(standard.New(":8000"))
}
