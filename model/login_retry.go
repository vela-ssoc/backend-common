package model

import "time"

type LoginRetry struct {
	CreatedAt time.Time `json:"created_at" json:"created_at"`
}

func (LoginRetry) TableName() string {
	return "login_retry"
}
