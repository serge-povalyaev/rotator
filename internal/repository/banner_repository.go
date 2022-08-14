package repository

import (
	"bannerRotator/internal/models"
	"github.com/jmoiron/sqlx"
)

type BannerRepository struct {
	conn *sqlx.DB
}

func NewBannerRepository(conn *sqlx.DB) *BannerRepository {
	return &BannerRepository{
		conn: conn,
	}
}

func (r *BannerRepository) Get(bannerID int) (*models.Banner, error) {
	var banner models.Banner
	sql := `SELECT * FROM banner WHERE banner_id = $1 LIMIT 1`
	err := r.conn.Get(&banner, sql, bannerID)
	if err != nil {
		return nil, err
	}

	return &banner, nil
}
