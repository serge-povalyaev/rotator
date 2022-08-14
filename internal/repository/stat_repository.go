package repository

import (
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	ActionTypeShow = iota
	ActionTypeClick
)

type StatRepository struct {
	conn *sqlx.DB
}

func NewStatRepository(conn *sqlx.DB) *StatRepository {
	return &StatRepository{
		conn: conn,
	}
}

func (r *StatRepository) Add(bannerID, slotID, socialGroupID, actionType int) error {
	sql := `
		INSERT INTO stat 
		(banner_id, slot_id, social_group_id, action_type, created_at) 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.conn.Exec(sql, bannerID, slotID, socialGroupID, actionType, time.Now())

	return err
}
