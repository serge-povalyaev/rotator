package repository

import (
	"bannerRotator/internal/models"
	"github.com/jmoiron/sqlx"
)

type SocialGroupRepository struct {
	conn *sqlx.DB
}

func NewSocialGroupRepository(conn *sqlx.DB) *SocialGroupRepository {
	return &SocialGroupRepository{
		conn: conn,
	}
}

func (r *SocialGroupRepository) Get(socialGroupID int) (*models.SocialGroup, error) {
	var socialGroup models.SocialGroup
	sql := `SELECT * FROM social_group WHERE social_group_id = $1 LIMIT 1`
	err := r.conn.Get(&socialGroup, sql, socialGroupID)
	if err != nil {
		return nil, err
	}

	return &socialGroup, nil
}
