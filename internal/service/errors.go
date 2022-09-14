package service

import "errors"

var (
	ErrBannerNotFound      = errors.New("баннер не найден")
	ErrSlotNotFound        = errors.New("слот не найден")
	ErrSocialGroupNotFound = errors.New("социальная группа не найдена")
)
