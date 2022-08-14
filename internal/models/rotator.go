package models

type BannerToSlot struct {
	BannerID int `db:"banner_id"`
	SlotID   int `db:"slot_id"`
}

type Banner struct {
	ID   int    `db:"banner_id"`
	Name string `db:"name"`
}

type Slot struct {
	ID   int    `db:"slot_id"`
	Name string `db:"name"`
}

type SocialGroup struct {
	ID   int    `db:"social_group_id"`
	Name string `db:"name"`
}

type Stat struct {
	BannerID      int    `db:"banner_id"`
	SlotID        int    `db:"slot_id"`
	SocialGroupID int    `db:"social_group_id"`
	Shows         int    `db:"shows"`
	Clicks        int    `db:"clicks"`
	UpdatedAt     string `db:"updated_at"`
}

func (s *Stat) GetID() int {
	return s.BannerID
}

func (s *Stat) GetTotalCount() int {
	return s.Shows
}

func (s *Stat) GetGoalsCount() int {
	return s.Clicks
}
