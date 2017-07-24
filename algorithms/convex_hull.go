package algorithms

import (
	"github.com/ClementDelgrange/gogeo"
	"math"
	"sort"
)


// Structure dataAngles is use to store values of angle between three folowwing coordinate in a list
// and id of the central coordinate
type dataAngles struct {
	idCoord int
	angle float64
}

type ByAngle []dataAngles

func (s ByAngle) Len() int {
	return len(s)
}
func (s ByAngle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByAngle) Less(i, j int) bool {
	return s[i].angle < s[j].angle
}


// Returns convex hull for a given coordinate list using Jarvis algorithm
func ConvexHullJarvis(coords []gogeo.Coordinate) []gogeo.Coordinate {
	var hull []gogeo.Coordinate
	leftId := FarLeftId(coords)
	hull = append(hull, coords[leftId])

	ini := leftId
	var fin int
	i := 1
	n := len(coords)

	for {
		fin = (ini + 1) % n
		for j := 2; j < len(coords); j++ {
			ccw := Counterclockwise(coords, ini, fin, (ini + j) % n)
			if ccw > 0 || ( ccw == 0 && FarAway(coords, ini, fin, (ini + j) % n) == (ini + j) % n ) {
				// coords[ini+j%n] est à gauche par rapport à la demi-droite [coord[ini],coords[fin])
				// ou sur la demi droite mais plus loin
				// coords[ini+j%n] est un bon candidat pour l'enveloppe : on met à jour
				fin = (ini + j) % n
			}
		}
		hull = append(hull, coords[fin])
		i++
		if fin == leftId {
			break
		} else {
			ini = fin
		}
	}

	return hull
}

// Returns convex hull for a given coordinate list using Jarvis algorithm
func ConvexHullGraham(coords []gogeo.Coordinate) []gogeo.Coordinate {
	var hull []gogeo.Coordinate
	leftId := FarLeftId(coords)
	n := len(coords)

	// calcule l'angle de coordonnée par rapport à la coordonnées la plus à gauche
	angles := make([]dataAngles, n)
	for i := 1; i <= n; i++ {
		angles[i-1].idCoord = (leftId + i) % n
		angles[i-1].angle = AngularAngleToLeftAwayPoint(coords, leftId, (leftId + i) % n)
	}

	// tri croissant des angles
	sort.Sort(sort.Reverse(ByAngle(angles)))

	// pour chaque triplet successif, déterminer s'il est dans l'enveloppe (triangle direct)
	var hullData []dataAngles
	hullData = append(hullData, dataAngles{leftId, 0})
	hullData = append(hullData, angles[0])
	var c1, c2 int

	for i := 1; i < n; i++ {
		for len(hullData) >= 2 {
			c1 = hullData[len(hullData)-2].idCoord
			c2 = hullData[len(hullData)-1].idCoord
			ccw := Counterclockwise(coords, c1, c2, angles[i].idCoord)
			if ccw > 0 {
				hullData = hullData[:len(hullData) - 1]
				continue
			} else if ccw == 0 {
				if FarAway(coords, hullData[len(hullData)-2].idCoord, hullData[len(hullData)-1].idCoord, angles[i].idCoord) == angles[i].idCoord {
					hullData[len(hullData) - 1] = angles[i]
				}
				break
			} else {
				hullData = append(hullData, angles[i])
				break
			}
		}
	}


	// construct hull
	for _, a := range hullData {
		hull = append(hull, coords[a.idCoord])
	}
	return hull
}

// Returns indice of the far away on the left coordinate
// In case of equality (several points of the same abscissa), returns the lowest point.
func FarLeftId(coords []gogeo.Coordinate) int {
	iLeft := 0
	for i := 1; i < len(coords); i++ {
		if coords[i][0] < coords[iLeft][0] {
			iLeft = i
		} else if coords[i][0] == coords[iLeft][0] && coords[i][1] < coords[iLeft][1] {
			iLeft = i
		}
	}
	return iLeft
}

// Determines if 3 coordinates are counterclockwise
// This function is base on the cross product.
// If result is positiv, coordinates are counterclockwise. If null, there are align. Else, there are clockwise.
func Counterclockwise(coords []gogeo.Coordinate, i, j, k int) float64 {
	res := (coords[j][0] - coords[i][0]) * (coords[k][1] - coords[i][1]) - (coords[j][1] - coords[i][1]) * (coords[k][0] - coords[i][0])
	return res
}

// Returns indice of the far away point from a given point
func FarAway(coords []gogeo.Coordinate, p int, i, j int) int {
	dpi := coords[p].DistanceTo(&coords[i])
	dpj := coords[p].DistanceTo(&coords[j])
	if dpi < dpj {
		return j
	} else {
		return i
	}
}

//
func AngularAngleToLeftAwayPoint(coords []gogeo.Coordinate, p int, i int) float64 {
	x := coords[i][0] - coords[p][0]
	y := coords[i][1] - coords[p][1]
	return math.Atan2(y, x)
}