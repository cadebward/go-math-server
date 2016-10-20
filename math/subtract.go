package math

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func SubtractHandler(ctx echo.Context) error {
	left, leftErr := strconv.Atoi(ctx.Param("left"))
	right, rightErr := strconv.Atoi(ctx.Param("right"))
	if leftErr != nil || rightErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	total := subtract(left, right)
	return ctx.String(http.StatusOK, strconv.Itoa(total))
}

func subtract(a int, b int) int {
	return a - b
}
