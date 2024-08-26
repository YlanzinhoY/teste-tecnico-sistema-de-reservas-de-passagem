package handler

import (
	"github.com/labstack/echo/v4"
	"time"
)

type HealthCheckerHandler struct {
	e *echo.Echo
}

func NewHealthCheckerHandler(echo *echo.Echo) *HealthCheckerHandler {
	return &HealthCheckerHandler{
		e: echo,
	}
}

func (s *HealthCheckerHandler) HealthChecker(c echo.Context) error {

	return c.JSON(200, map[string]interface{}{
		"status": "ok",
		"time":   time.Now(),
	})

}
