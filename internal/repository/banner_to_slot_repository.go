package repository

import (
	"database/sql"
	"errors"

	"bannerRotator/internal/models"
	"github.com/jmoiron/sqlx"
)

type BannerToSlotRepository struct {
	conn *sqlx.DB
}

func NewBannerToSlotRepository(conn *sqlx.DB) *BannerToSlotRepository {
	return &BannerToSlotRepository{
		conn: conn,
	}
}

func (r *BannerToSlotRepository) AddBannerToSlot(bannerID, slotID int) error {
	exists, err := r.existsBannerToSlot(bannerID, slotID)
	if err != nil {
		return err
	}

	if exists {
		return ErrBannerToSlotExists
	}

	query := `INSERT INTO banner_to_slot (banner_id, slot_id) VALUES ($1, $2)`
	_, err = r.conn.Exec(query, bannerID, slotID)

	if err != nil {
		return err
	}

	return nil
}

func (r *BannerToSlotRepository) RemoveBannerToSlot(bannerID, slotID int) error {
	exists, err := r.existsBannerToSlot(bannerID, slotID)
	if err != nil {
		return err
	}

	if !exists {
		return ErrBannerToSlotNotExists
	}

	query := `DELETE FROM banner_to_slot WHERE banner_id = $1 AND slot_id = $2`
	_, err = r.conn.Exec(query, bannerID, slotID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BannerToSlotRepository) existsBannerToSlot(bannerID, slotID int) (bool, error) {
	var bannerToSlot models.BannerToSlot
	query := `SELECT * FROM banner_to_slot WHERE banner_id = $1 AND slot_id = $2 LIMIT 1`
	err := r.conn.Get(&bannerToSlot, query, bannerID, slotID)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *BannerToSlotRepository) GetBanners(slotID int) ([]models.BannerToSlot, error) {
	var bannerToSlot []models.BannerToSlot
	query := `SELECT * FROM banner_to_slot WHERE slot_id = $1`
	err := r.conn.Select(&bannerToSlot, query, slotID)

	if errors.Is(err, sql.ErrNoRows) || err == nil {
		return bannerToSlot, nil
	}

	return nil, err
}
