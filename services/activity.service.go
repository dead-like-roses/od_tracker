package services

import (
	"database/sql"
	"time"

	"github.com/dead-like-roses/od_tracker/.gen/od_track/public/model"
	. "github.com/dead-like-roses/od_tracker/.gen/od_track/public/table"
	. "github.com/go-jet/jet/postgres"
)

func NewActivityService(db *sql.DB) *ActivityService {
	return &ActivityService{
		db,
	}
}

type ActivityService struct {
	db *sql.DB
}

func (as ActivityService) CreateDataPoint(device string, posted time.Time) {
	data := model.Activity{DeviceID: device, PostedAt: posted, App: "none"}

	insertStmt := INSERT(Activity.DeviceID, Activity.PostedAt, Activity.App).MODEL(data)
	insertStmt.Exec(as.db)
}

func (as ActivityService) CreateMultipleDataPoints(device string, times []time.Time) {
	dataPoints := make([]model.Activity, len(times))
	for i, t := range times {
		dataPoints[i] = model.Activity{DeviceID: device, PostedAt: t, App: "none"}
	}

	insertStmt := INSERT(Activity.DeviceID, Activity.PostedAt, Activity.App).MODELS(dataPoints)
	insertStmt.Exec(as.db)
}

func (as ActivityService) ListAllDatapoints(device string) []model.Activity {
	dataPoints := SELECT(
		Activity.DeviceID,
		Activity.PostedAt,
		Activity.App,
	).FROM(
		Activity,
	).WHERE(
		Activity.DeviceID.EQ(String(device)),
	)
}
