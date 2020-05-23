package services

import (
	"sort"

	"github.com/aymone/cartesian-api/models"
	"github.com/aymone/cartesian-api/utils"
)

// PointsService interface
type PointsService interface {
	GetPoints(x, y, distance int) []*models.Points
}

type pointsService struct {
	points []*models.Points
}

// NewPointsService initialize a points service
func NewPointsService(p []*models.Points) *pointsService {
	return &pointsService{
		points: p,
	}
}

func (s *pointsService) GetPoints(x, y, distance int) []*models.Points {
	var list []*models.Points

	p := &models.Points{
		X:        x,
		Y:        y,
		Distance: distance,
	}

	for _, v := range s.points {
		distance := calculate(p, v)
		if p.Distance >= distance {
			list = append(list, &models.Points{
				X:        v.X,
				Y:        v.Y,
				Distance: distance,
			})
		}
	}

	sort.Sort(models.DistanceSorter(list))

	return list
}

// Calculate returns the distance between two points
func calculate(a, b *models.Points) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}
