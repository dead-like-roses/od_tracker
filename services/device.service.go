package services

import (
	"database/sql"
	// . "github.com/go-jet/jet/v2/postgres"
)

func NewDeviceService(db *sql.DB) *DeviceService {
	return &DeviceService{
		db,
	}
}

type DeviceService struct {
	db *sql.DB
}

func (ds DeviceService) registerDevice(name string) {
}
