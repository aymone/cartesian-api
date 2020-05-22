package services

import (
	"testing"

	"github.com/aymone/cartesian-api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetPoints(t *testing.T) {
	data := []*models.Points{
		{X: 5, Y: 2},
		{X: 4, Y: 2},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 13, Y: 11},
	}

	s := NewPointsService(data)

	p := models.Points{
		Distance: 5,
		X:        3,
		Y:        1,
	}

	list := s.GetPoints(&p)

	// evaluate quantity of results
	assert.Len(t, list, 4)

	// evaluate order of results
	assert.Equal(t, list[0].Distance, 0)
	assert.Equal(t, list[1].Distance, 1)
	assert.Equal(t, list[2].Distance, 2)
	assert.Equal(t, list[3].Distance, 3)
}
