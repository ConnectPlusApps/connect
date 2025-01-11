package models

import (
    "time"
)

type Profile struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"uniqueIndex;not null"`
    DisplayName string    `gorm:"size:100;not null"`
    Bio         string    `gorm:"size:500"`
    Gender      string    `gorm:"size:50"`
    BirthDate   time.Time
    Location    string    `gorm:"size:100"`
    Photos      []string  `gorm:"serializer:json"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
