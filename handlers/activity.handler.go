package handlers

import (
	"net/http"
	"time"

	"github.com/dead-like-roses/od_tracker/services"
	"github.com/labstack/echo/v4"
)

type ActivityRequest struct {
	Device   string    `json:"device"`
	PostedAt time.Time `json:"posted_at"`
}

type ActivityHandler struct {
	service services.ActivityService
}

func (ah ActivityHandler) registerActivity(c echo.Context) error {
	a := new(ActivityRequest)
	if err := c.Bind(a); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, a)
}

func (ah ActivityHandler) listDatapoints(c echo.Context) error {
	return c.String(http.StatusOK, "list of datapoints")
}
