package services

import (
	"database/sql"

	"github.com/dead-like-roses/od_tracker/.gen/od_track/public/model"
	"github.com/dead-like-roses/od_tracker/.gen/od_track/public/table"
	"github.com/google/uuid"
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

func (ds DeviceService) RegisterDevice(name string) string {
	deviceID := uuid.New().String()

	device := model.Devices{
		DeviceID: deviceID,
		Name:     name,
	}

	insertStmt := table.Devices.INSERT(table.Devices.DeviceID, table.Devices.Name).MODEL(device)
	insertStmt.Exec(ds.db)
	return deviceID
}
