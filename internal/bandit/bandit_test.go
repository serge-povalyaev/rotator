package bandit

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type EntityTest struct {
	ID         int
	TotalCount int
	GoalsCount int
}

func (s *EntityTest) GetID() int {
	return s.ID
}

func (s *EntityTest) GetTotalCount() int {
	return s.TotalCount
}

func (s *EntityTest) GetGoalsCount() int {
	return s.GoalsCount
}

func TestBandit(t *testing.T) {
	t.Run("Пустой список для сравнения", func(t *testing.T) {
		bandit := NewBandit([]Entity{})
		element, err := bandit.SelectElement()
		require.Nil(t, element)
		require.ErrorIs(t, err, ErrEmptyEntities)
	})

	t.Run("Список с элементами с нулевым TotalCount", func(t *testing.T) {
		bandit := NewBandit([]Entity{
			&EntityTest{
				ID:         1,
				TotalCount: 1,
				GoalsCount: 1,
			},
			&EntityTest{
				ID:         2,
				TotalCount: 2,
				GoalsCount: 2,
			},
			&EntityTest{
				ID:         3,
				TotalCount: 0,
				GoalsCount: 0,
			},
		})
		element, err := bandit.SelectElement()
		require.NoError(t, err)
		require.Equal(t, 3, element.GetID())
	})

	t.Run("Выбор элемента по формуле", func(t *testing.T) {
		bandit := NewBandit([]Entity{
			&EntityTest{
				ID:         1,
				TotalCount: 4,
				GoalsCount: 1,
			},
			&EntityTest{
				ID:         2,
				TotalCount: 2,
				GoalsCount: 0,
			},
			&EntityTest{
				ID:         3,
				TotalCount: 2,
				GoalsCount: 2,
			},
			&EntityTest{
				ID:         4,
				TotalCount: 2,
				GoalsCount: 2,
			},
			&EntityTest{
				ID:         5,
				TotalCount: 1,
				GoalsCount: 0,
			},
		})
		element, err := bandit.SelectElement()
		require.NoError(t, err)
		require.Equal(t, 4, element.GetID())
	})
}
