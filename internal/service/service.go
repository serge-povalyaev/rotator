package service

import (
	"bannerRotator/internal/bandit"
	"bannerRotator/internal/logger"
	"bannerRotator/internal/models"
	"bannerRotator/internal/rabbit"
	"bannerRotator/internal/repository"
	"encoding/json"
	"time"
)

type RotatorService struct {
	logger                 *logger.Logger
	bannerRepository       *repository.BannerRepository
	bannerToSlotRepository *repository.BannerToSlotRepository
	slotRepository         *repository.SlotRepository
	socialGroupRepository  *repository.SocialGroupRepository
	statRepository         *repository.StatRepository
	totalStatRepository    *repository.TotalStatRepository
	producer               *rabbit.Producer
}

func NewRotatorService(
	logger *logger.Logger,
	bannerRepository *repository.BannerRepository,
	bannerToSlotRepository *repository.BannerToSlotRepository,
	slotRepository *repository.SlotRepository,
	socialGroupRepository *repository.SocialGroupRepository,
	statRepository *repository.StatRepository,
	totalStatRepository *repository.TotalStatRepository,
	producer *rabbit.Producer,
) *RotatorService {
	return &RotatorService{
		logger:                 logger,
		bannerRepository:       bannerRepository,
		bannerToSlotRepository: bannerToSlotRepository,
		slotRepository:         slotRepository,
		socialGroupRepository:  socialGroupRepository,
		statRepository:         statRepository,
		totalStatRepository:    totalStatRepository,
		producer:               producer,
	}
}

func (s *RotatorService) AddBanner(bannerID, slotID int) error {
	if err := s.checkBannerAndSlot(bannerID, slotID); err != nil {
		return err
	}

	return s.bannerToSlotRepository.AddBannerToSlot(bannerID, slotID)
}

func (s *RotatorService) RemoveBanner(bannerID, slotID int) error {
	if err := s.checkBannerAndSlot(bannerID, slotID); err != nil {
		return err
	}

	return s.bannerToSlotRepository.RemoveBannerToSlot(bannerID, slotID)
}

func (s *RotatorService) ClickBanner(bannerID, slotID, socialGroupID int) error {
	if err := s.checkBannerAndSlot(bannerID, slotID); err != nil {
		return err
	}

	if err := s.checkSocialGroup(socialGroupID); err != nil {
		return err
	}

	err := s.statRepository.Add(bannerID, slotID, socialGroupID, repository.ActionTypeClick)
	if err != nil {
		return err
	}

	err = s.totalStatRepository.IncrementClicks(bannerID, slotID, socialGroupID)
	if err != nil {
		return err
	}

	return s.sendEvent(slotID, bannerID, socialGroupID, models.EventTypeClick)
}

func (s *RotatorService) Get(slotID, socialGroupID int) (*int, error) {
	if err := s.checkSlot(slotID); err != nil {
		return nil, err
	}

	if err := s.checkSocialGroup(socialGroupID); err != nil {
		return nil, err
	}

	selectedEntity, err := s.getBanner(slotID, socialGroupID)
	if err != nil {
		return nil, err
	}

	bannerID := selectedEntity.GetID()

	err = s.statRepository.Add(bannerID, slotID, socialGroupID, repository.ActionTypeShow)
	if err != nil {
		return nil, err
	}

	err = s.totalStatRepository.IncrementShows(bannerID, slotID, socialGroupID)
	if err != nil {
		return nil, err
	}

	err = s.sendEvent(slotID, bannerID, socialGroupID, models.EventTypeShow)
	if err != nil {
		return nil, err
	}

	return &bannerID, nil
}

func (s *RotatorService) getBanner(slotID, socialGroupID int) (bandit.Entity, error) {
	stat, err := s.totalStatRepository.GetStat(slotID, socialGroupID)
	if err != nil {
		return nil, err
	}

	banners, err := s.bannerToSlotRepository.GetBanners(slotID)
	if err != nil {
		return nil, err
	}

	entitiesMap := make(map[int]bandit.Entity, len(stat))
	for _, item := range stat {
		item := item
		entitiesMap[item.BannerID] = bandit.Entity(&item)
	}

	for _, banner := range banners {
		_, ok := entitiesMap[banner.BannerID]
		if ok == true {
			continue
		}

		entitiesMap[banner.BannerID] = bandit.Entity(&models.Stat{BannerID: banner.BannerID})
	}

	entities := make([]bandit.Entity, 0, len(entitiesMap))
	for _, entity := range entitiesMap {
		entities = append(entities, entity)
	}

	banditModel := bandit.NewBandit(entities)
	selectedEntity, err := banditModel.SelectElement()
	if err != nil {
		return nil, err
	}

	return selectedEntity, nil
}

func (s *RotatorService) sendEvent(slotID, bannerID, socialGroupID, eventType int) error {
	message := &models.Event{
		Type:          eventType,
		SlotID:        slotID,
		BannerID:      bannerID,
		SocialGroupID: socialGroupID,
		DateTime:      time.Now(),
	}

	messageBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return s.producer.Publish(messageBody)
}

func (s *RotatorService) checkBanner(bannerID int) error {
	banner, err := s.bannerRepository.Get(bannerID)
	if err != nil {
		return err
	}

	if banner == nil {
		return ErrBannerNotFound
	}

	return nil
}

func (s *RotatorService) checkSlot(slotID int) error {
	slot, err := s.slotRepository.Get(slotID)
	if err != nil {
		return err
	}

	if slot == nil {
		return ErrSlotNotFound
	}

	return nil
}

func (s *RotatorService) checkSocialGroup(socialGroupID int) error {
	socialGroup, err := s.socialGroupRepository.Get(socialGroupID)
	if err != nil {
		return err
	}

	if socialGroup == nil {
		return ErrSocialGroupNotFound
	}

	return nil
}

func (s *RotatorService) checkBannerAndSlot(bannerID, slotID int) error {
	err := s.checkBanner(bannerID)
	if err == nil {
		err = s.checkSlot(slotID)
	}

	return err
}
