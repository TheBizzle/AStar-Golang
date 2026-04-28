// Package heuristic defines the two common distance heuristics: Euclidean distance (AKA as the crow flies,
// AKA the hypotenuse of a triangle) and Manhattan distance (AKA taxicab distance, AKA the sum of the two
// legs of a right triangle).
package heuristic

import (
	"fmt"
	"math"

	core "github.com/TheBizzle/PathFindingCore-Golang/pathingmap"
)

//go:generate stringer -type=Heuristic
type Heuristic uint8

const (
	Euclidean Heuristic = iota
	Manhattan
	Dijkstra
)

func (h Heuristic) Distance(c1, c2 core.Coordinate) float64 {
	switch h {
	case Euclidean:
		xComp := sub(c1.X, c2.X)
		yComp := sub(c1.Y, c2.Y)
		comps := (xComp * xComp) + (yComp * yComp)
		return math.Sqrt(float64(comps))
	case Manhattan:
		sum := sub(c1.X, c2.X) + sub(c1.Y, c2.Y)
		return float64(sum)
	case Dijkstra:
		return 0
	default:
		panic(fmt.Sprintf("Impossible heuristic (Distance): %v", h))
	}
}

func sub(u1, u2 uint) uint {
	if u1 >= u2 {
		return u1 - u2
	} else {
		return u2 - u1
	}
}
