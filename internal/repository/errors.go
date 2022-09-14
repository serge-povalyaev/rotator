package repository

import "errors"

var (
	ErrBannerToSlotExists    = errors.New("баннер уже присутствует в указанном слоте")
	ErrBannerToSlotNotExists = errors.New("баннер не присутствует в указанном слоте")
	ErrStatNotExists         = errors.New("статистика не существует")
)
