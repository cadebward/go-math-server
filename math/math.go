package math

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// AddHandler add two numbers handler
func AddHandler(ctx echo.Context) error {
	left, leftErr := strconv.Atoi(ctx.Param("left"))
	if leftErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	right, rightErr := strconv.Atoi(ctx.Param("right"))
	if rightErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	total := left + right
	return ctx.String(http.StatusOK, strconv.Itoa(total))
}
