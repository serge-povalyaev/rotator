package bandit

import (
	"math"
)

type Entity interface {
	GetID() int
	GetTotalCount() int
	GetGoalsCount() int
}

type Bandit struct {
	Entities []Entity
}

func NewBandit(entities []Entity) *Bandit {
	return &Bandit{
		Entities: entities,
	}
}

func (b *Bandit) SelectElement() (Entity, error) {
	if len(b.Entities) == 0 {
		return nil, ErrEmptyEntities
	}

	entity := b.getWithEmptyTotal()
	if entity != nil {
		return entity, nil
	}

	totalCount := b.getTotal()

	var maxPriority float64
	var selectedEntity Entity
	for _, entity = range b.Entities {
		partGoals := float64(entity.GetGoalsCount()) / float64(entity.GetTotalCount())
		entityPriority := partGoals + math.Sqrt(2*math.Log(float64(totalCount))/float64(entity.GetTotalCount()))
		if entityPriority >= maxPriority {
			maxPriority = entityPriority
			selectedEntity = entity
		}
	}

	return selectedEntity, nil
}

func (b *Bandit) getWithEmptyTotal() Entity {
	for _, entity := range b.Entities {
		if entity.GetTotalCount() == 0 {
			return entity
		}
	}

	return nil
}

func (b *Bandit) getTotal() int {
	var totalCount int
	for _, entity := range b.Entities {
		totalCount += entity.GetTotalCount()
	}

	return totalCount
}
