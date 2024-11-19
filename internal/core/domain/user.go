package domain

import "time"

type User struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"unique"`
	Password  string
	Status    bool `gorm:"not null;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
