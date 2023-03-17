package model

import "time"

type JobMinion struct {
	ID        int64
	JobID     int64
	MinionID  int64
	Inet      string
	Success   bool
	Status    uint8
	Cause     string
	CreatedAt time.Time
}

func (JobMinion) TableName() string {
	return "job_minion"
}
