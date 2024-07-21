package services

import (
	"database/sql"
	"time"

	"github.com/dead-like-roses/od_tracker/.gen/od_track/public/model"
	"github.com/dead-like-roses/od_tracker/.gen/od_track/public/table"
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
	data := model.Activity{DeviceID: device, PostedAt: posted, App: "none"}

	insertStmt := table.Activity.INSERT(table.Activity.DeviceID, table.Activity.PostedAt, table.Activity.App).MODEL(data)
	insertStmt.Exec(as.db)
}

func (as ActivityService) createMultipleDataPoints(device string, times []time.Time) {
	dataPoints := make([]model.Activity, len(times))
	for i, t := range times {
		dataPoints[i] = model.Activity{DeviceID: device, PostedAt: t, App: "none"}
	}

	insertStmt := table.Activity.INSERT(table.Activity.DeviceID, table.Activity.PostedAt, table.Activity.App).MODELS(dataPoints)
	insertStmt.Exec(as.db)
}
