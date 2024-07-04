package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ActivityDatapoint struct {
	DeviceID int       `json:"device_id"`
	PostedAt time.Time `json:"posted_at"`
}

type ActivityHandler struct {
}

func (ah ActivityHandler) registerActivity(c echo.Context) error {
	return c.String(http.StatusCreated, "registered datapoint")
}

func main() {
	e := echo.New()

	ah := ActivityHandler{}

	e.POST("/register", ah.registerActivity)

	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
