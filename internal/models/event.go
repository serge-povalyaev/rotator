package models

import "time"

const (
	EventTypeShow = iota
	EventTypeClick
)

type Event struct {
	Type          int       `json:"type"`
	SlotID        int       `json:"slotId"`
	BannerID      int       `json:"bannerId"`
	SocialGroupID int       `json:"socialGroupId"`
	DateTime      time.Time `json:"datetime"`
}
