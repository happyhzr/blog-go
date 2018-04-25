package model

import "time"

type Model struct {
	ID          int64      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	CreatedAtTS int64      `gorm:"-" json:"created_at"`
	UpdatedAtTS int64      `gorm:"-" json:"updated_at"`
}
