package model

import "time"

type JobReport struct {
	ID        int64
	JobID     int64
	MinionID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName implemented schema.Tabler
func (JobReport) TableName() string {
	return "job_report"
}
