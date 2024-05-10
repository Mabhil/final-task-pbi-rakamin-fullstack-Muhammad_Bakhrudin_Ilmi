package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Photos    []Photo   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"photos"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
