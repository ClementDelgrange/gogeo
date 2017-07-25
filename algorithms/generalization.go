package algorithms

import (
	"github.com/ClementDelgrange/gogeo"
)

// Simplifies a linestring using Douglas-Peuker algorithm.
// See : https://en.wikipedia.org/wiki/Ramer–Douglas–Peucker_algorithm.
func DouglasPeucker(coords []gogeo.Coordinate, e float64) []gogeo.Coordinate {
	n := len(coords)
	if n <= 2 {
		return coords
	}

	var dmax float64 = -1
	var imax int = -1
	for i := 1; i < n-1; i++ {
		d, _ := coords[i].ShortestDistanceToSegment(&coords[0], &coords[n-1])
		if d > dmax {
			dmax = d
			imax = i
		}
	}

	if dmax < e {
		return []gogeo.Coordinate{coords[0], coords[n-1]}
	} else {
		c1 := DouglasPeucker(coords[:imax+1], e)
		c2 := DouglasPeucker(coords[imax:], e)
		return append(c1, c2[1:]...)
	}
}




