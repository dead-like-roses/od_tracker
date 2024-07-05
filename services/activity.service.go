package services

import (
	"database/sql"
	"time"
)

func NewActivityService(db *sql.DB) *ActivityService {
	return &ActivityService{
		db,
	}
}

type ActivityService struct {
	db *sql.DB
}

func (as ActivityService) createDataPoint(device string, posted time.Time) {
}
