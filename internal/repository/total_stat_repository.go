package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"bannerRotator/internal/models"
	"github.com/jmoiron/sqlx"
)

type TotalStatRepository struct {
	conn *sqlx.DB
}

func NewTotalStatRepository(conn *sqlx.DB) *TotalStatRepository {
	return &TotalStatRepository{
		conn: conn,
	}
}

func (r *TotalStatRepository) GetStat(slotID, socialGroupID int) ([]models.Stat, error) {
	var stat []models.Stat
	query := `SELECT * FROM total_stat WHERE slot_id = $1 AND social_group_id = $2`
	err := r.conn.Select(&stat, query, slotID, socialGroupID)

	if errors.Is(err, sql.ErrNoRows) {
		return []models.Stat{}, nil
	}

	if err != nil {
		return nil, err
	}

	return stat, nil
}

func (r *TotalStatRepository) IncrementShows(bannerID, slotID, socialGroupID int) error {
	err := r.incrementStat(bannerID, slotID, socialGroupID, "shows")
	if !errors.Is(err, ErrStatNotExists) {
		return err
	}

	query := `
		INSERT INTO total_stat
		(banner_id, slot_id, social_group_id, shows, clicks, updated_at) 
		VALUES 
		($1, $2, $3, 1, 0, $4)
	`

	_, err = r.conn.Exec(query, bannerID, slotID, socialGroupID, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *TotalStatRepository) IncrementClicks(bannerID, slotID, socialGroupID int) error {
	return r.incrementStat(bannerID, slotID, socialGroupID, "clicks")
}

func (r *TotalStatRepository) incrementStat(bannerID, slotID, socialGroupID int, column string) error {
	query := fmt.Sprintf(`
		UPDATE total_stat 
		SET %s = %s + 1, updated_at = $1 
		WHERE banner_id = $2 AND slot_id = $3 AND social_group_id = $4
	`, column, column)
	result, err := r.conn.Exec(query, time.Now(), bannerID, slotID, socialGroupID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 1 {
		return nil
	}

	return ErrStatNotExists
}
