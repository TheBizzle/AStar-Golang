// Package stepdata defines the core data structures for representing the state of a pathfinding attempt.
package stepdata

import (
	cq "github.com/TheBizzle/AStar-Golang/coordqueue"
	heur "github.com/TheBizzle/AStar-Golang/heuristic"
	opt "github.com/TheBizzle/AStar-Golang/option"
	core "github.com/TheBizzle/PathFindingCore-Golang/pathingmap"
)

type LocationData struct {
	Breadcrumb core.Breadcrumb
	CostOpt    opt.Option[float64]
	WasVisited bool
}

type StepData struct {
	LocDataMap   map[core.Coordinate]LocationData
	MaxIters     uint
	PathingMap   core.PathingMap
	GoalCoord    core.Coordinate
	CurrentCoord core.Coordinate
	ItersSoFar   uint
	Heuristic    heur.Heuristic
	Queue        cq.CoordQueue
}
