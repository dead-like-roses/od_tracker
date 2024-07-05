package handlers

import (
	"net/http"

	"github.com/dead-like-roses/od_tracker/services"
	"github.com/labstack/echo/v4"
)

type ActivityHandler struct {
	service services.ActivityService
}

func (ah ActivityHandler) registerActivity(c echo.Context) error {
	return c.String(http.StatusCreated, "registered datapoint")
}

func (ah ActivityHandler) listDatapoints(c echo.Context) error {
	return c.String(http.StatusOK, "list of datapoints")
}
