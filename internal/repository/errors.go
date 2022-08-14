package repository

import "errors"

var ErrBannerToSlotExists = errors.New("баннер уже присутствует в указанном слоте")
var ErrBannerToSlotNotExists = errors.New("баннер не присутствует в указанном слоте")
var ErrStatNotExists = errors.New("статистика не существует")
