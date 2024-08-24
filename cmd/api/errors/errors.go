package errors

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IdNotFound(err error, echo echo.Context) error {
	if errors.Is(err, sql.ErrNoRows) {
		return echo.JSON(http.StatusNotFound, map[string]string{
			"error": "id not found",
		})
	}

	return nil
}
