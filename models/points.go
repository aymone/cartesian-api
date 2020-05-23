package models

// Points represents two coordinates
type Points struct {
	X        int `json:"x"`
	Y        int `json:"y"`
	Distance int `json:"-"`
}

// DistanceSorter implements sort.Interface based on the Points.Distance field.
type DistanceSorter []*Points

func (d DistanceSorter) Len() int { return len(d) }

func (d DistanceSorter) Less(i, j int) bool { return d[i].Distance < d[j].Distance }

func (d DistanceSorter) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
