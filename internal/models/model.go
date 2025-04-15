package models

import (
	"time"

	"github.com/lib/pq"
)

type Campaign struct {
	ID        string         `json:"id"`
	Budget    float64        `json:"budget"`
	ReachGoal int            `json:"reach_goal"`
	Platforms pq.StringArray `gorm:"type:text[]" json:"platforms"`
}

type Bid struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CampaignID   string    `json:"campaign_id"`
	Platform     string    `json:"platform"`
	Amount       float64   `json:"amount"`
	EstimatedCVR float64   `json:"estimated_cvr"` // 👈 new field
	CreatedAt    time.Time `json:"created_at"`
}
