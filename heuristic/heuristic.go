// Package heuristic defines the heuristics for use in the A* pathfinding algorithm.  One of them (Dijkstra)
// is a do-nothing heuristic that turns A* into Dijkstra's (less-efficient) algorithm.  The others are
// common distance heuristics: Euclidean distance (AKA as the crow flies, AKA the hypotenuse of a triangle)
// and Manhattan distance (AKA taxicab distance, AKA the sum of the two legs of a right triangle).
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

func (h Heuristic) Eval(c1, c2 core.Coordinate) float64 {
	switch h {
	case Euclidean:
		xComp := difference(c1.X, c2.X)
		yComp := difference(c1.Y, c2.Y)
		comps := (xComp * xComp) + (yComp * yComp)
		return math.Sqrt(float64(comps))
	case Manhattan:
		sum := difference(c1.X, c2.X) + difference(c1.Y, c2.Y)
		return float64(sum)
	case Dijkstra:
		return 0
	default:
		panic(fmt.Sprintf("Impossible heuristic: %v", h))
	}
}

func difference(u1, u2 uint) uint {
	if u1 >= u2 {
		return u1 - u2
	} else {
		return u2 - u1
	}
}
