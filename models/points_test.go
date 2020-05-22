package models

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorterByDistance(t *testing.T) {
	list := []*Points{
		{Distance: 4},
		{Distance: 3},
		{Distance: 1},
		{Distance: 4},
	}

	sort.Sort(DistanceSorter(list))

	assert.Equal(t, list[0].Distance, 1)
	assert.Equal(t, list[1].Distance, 3)
	assert.Equal(t, list[2].Distance, 4)
	assert.Equal(t, list[3].Distance, 4)
}
