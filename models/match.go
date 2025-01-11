package models

import (
    "time"
)

type MatchStatus string

const (
    MatchPending  MatchStatus = "pending"
    MatchAccepted MatchStatus = "accepted"
    MatchDeclined MatchStatus = "declined"
)

type Match struct {
    ID          uint       `gorm:"primaryKey"`
    User1ID     uint       `gorm:"not null"`
    User2ID     uint       `gorm:"not null"`
    Status      MatchStatus `gorm:"type:varchar(20);default:'pending'"`
    CreatedAt   time.Time  `gorm:"autoCreateTime"`
    UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}
