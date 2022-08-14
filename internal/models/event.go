package models

import "time"

const (
	EventTypeShow = iota
	EventTypeClick
)

type Event struct {
	Type          int       `json:"type"`
	SlotID        int       `json:"slot_id"`
	BannerID      int       `json:"banner_id"`
	SocialGroupID int       `json:"social_group_id"`
	DateTime      time.Time `json:"datetime"`
}
