// Package coordqueue defines a coordinate queue data structure, as a Golang heap.  In A*, the coordinate
// queue tells us which coordinate is the highest priority for us to check next, as determined by the
// distance heuristic.
package coordqueue

import (
	core "github.com/TheBizzle/PathFindingCore-Golang/pathingmap"
)

type PrioBundle[T any] struct {
	Priority float64
	Item     T
}

type MiniLoc struct {
	Breadcrumb core.Breadcrumb
	Coord      core.Coordinate
	Cost       float64
}

type CoordQueue []PrioBundle[MiniLoc]

func (q CoordQueue) Len() int {
	return len(q)
}

func (q CoordQueue) Less(i, j int) bool {
	return q[i].Priority < q[j].Priority
}

func (q CoordQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *CoordQueue) Push(x any) {
	*q = append(*q, x.(PrioBundle[MiniLoc]))
}

func (q *CoordQueue) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
