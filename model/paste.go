package model

import "time"

type Paste struct {
	ID       string    `json:"id" gorm:"primaryKey"`
	ExpireAt time.Time `json:"expire"`
	Type     string    `json:"type"`
	Data     string    `json:"data"`
}
