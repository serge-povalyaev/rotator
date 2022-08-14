package repository

import (
	"bannerRotator/internal/models"
	"github.com/jmoiron/sqlx"
)

type SlotRepository struct {
	conn *sqlx.DB
}

func NewSlotRepository(conn *sqlx.DB) *SlotRepository {
	return &SlotRepository{
		conn: conn,
	}
}

func (r *SlotRepository) Get(slotID int) (*models.Slot, error) {
	var slot models.Slot
	sql := `SELECT * FROM slot WHERE slot_id = $1 LIMIT 1`
	err := r.conn.Get(&slot, sql, slotID)
	if err != nil {
		return nil, err
	}

	return &slot, nil
}
