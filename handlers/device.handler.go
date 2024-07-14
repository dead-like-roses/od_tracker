package handlers

import (
	"net/http"

	"github.com/dead-like-roses/od_tracker/services"
	"github.com/labstack/echo/v4"
)

type DeviceRequest struct {
	Name string `json:"name"`
}

type DeviceResponse struct {
	Name     string `json:"name"`
	DeviceID string `json:"device_id"`
}

func NewDeviceHandler(service *services.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		service,
	}
}

type DeviceHandler struct {
	service *services.DeviceService
}

func (dh DeviceHandler) RegisterDevice(c echo.Context) error {
	d := new(DeviceRequest)
	if err := c.Bind(d); err != nil {
		return err
	}

	id := dh.service.RegisterDevice(d.Name)

	res := DeviceResponse{
		Name:     d.Name,
		DeviceID: id,
	}

	return c.JSON(http.StatusCreated, res)
}
