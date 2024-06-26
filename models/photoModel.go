package models

import "time"

type Photo struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Caption   string    `gorm:"type:varchar(255);not null" json:"caption"`
	PhotoUrl  string    `gorm:"type:varchar(255);not null" json:"photo_url"`
	UserID    uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
