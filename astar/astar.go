// Package astar defines the algorithm for unidirectional A* pathfinding.
package astar

import (
	"container/heap"

	cq "github.com/TheBizzle/AStar-Golang/coordqueue"
	heur "github.com/TheBizzle/AStar-Golang/heuristic"
	sd "github.com/TheBizzle/AStar-Golang/stepdata"
	core "github.com/TheBizzle/PathFindingCore-Golang/pathingmap"
)

func Run(pms core.PathingMapString, heuristic heur.Heuristic) (core.RunResult, sd.StepData) {
	stepData := genInitialState(pms, heuristic)
	for {
		isPrimed := primeNextStep(&stepData)
		if stepData.GoalCoord == stepData.CurrentCoord {
			return core.SuccessfulRun, stepData
		} else if !isPrimed || stepData.ItersSoFar >= stepData.MaxIters {
			return core.FailedRun, stepData
		} else {
			for _, neighbor := range stepData.PathingMap.NeighborsOf(stepData.CurrentCoord) {
				if !stepData.LocDataMap[neighbor].WasVisited {
					enqueueNeighbor(neighbor, &stepData)
				}
			}
			locData := stepData.LocDataMap[stepData.CurrentCoord]
			locData.WasVisited = true
			stepData.LocDataMap[stepData.CurrentCoord] = locData
			stepData.ItersSoFar++
		}
	}
}

func genInitialState(pms core.PathingMapString, heuristic heur.Heuristic) sd.StepData {
	pmd := pms.AsPMD()
	pmap := core.PathingMap{Grid: pmd.Grid}

	selfBreadcrumb := core.Source{Coord: pmd.Start}
	selfLoc := cq.MiniLoc{Breadcrumb: selfBreadcrumb, Coord: pmd.Start, Cost: 0}

	queue := cq.CoordQueue{cq.PrioBundle[cq.MiniLoc]{Priority: 0, Item: selfLoc}}
	heap.Init(&queue)

	locData := map[core.Coordinate]sd.LocationData{}

	for coord := range pmd.Grid {
		locData[coord] = sd.LocationData{Breadcrumb: nil, Cost: nil, WasVisited: false}
	}

	locData[pmd.Start] = sd.LocationData{Breadcrumb: selfBreadcrumb, Cost: new(0.0), WasVisited: false}

	return sd.StepData{
		LocDataMap:   locData,
		MaxIters:     pmap.Height() * pmap.Width(),
		PathingMap:   pmap,
		GoalCoord:    pmd.Goal,
		CurrentCoord: pmd.Start,
		ItersSoFar:   0,
		Heuristic:    heuristic,
		Queue:        queue,
	}
}

func primeNextStep(stepData *sd.StepData) bool {
	var next cq.PrioBundle[cq.MiniLoc]

	for {
		if stepData.Queue.Len() == 0 {
			return false
		}
		next = heap.Pop(&stepData.Queue).(cq.PrioBundle[cq.MiniLoc])
		if !stepData.LocDataMap[next.Item.Coord].WasVisited {
			break
		}
	}

	nextItem := next.Item

	locData := stepData.LocDataMap[nextItem.Coord]
	locData.Breadcrumb = nextItem.Breadcrumb
	locData.Cost = &nextItem.Cost
	stepData.LocDataMap[nextItem.Coord] = locData
	stepData.CurrentCoord = nextItem.Coord

	return true
}

func enqueueNeighbor(neighbor core.Coordinate, stepData *sd.StepData) {
	currentLoc := stepData.LocDataMap[stepData.CurrentCoord]
	newCost := *currentLoc.Cost + 1
	if stepData.LocDataMap[neighbor].Cost == nil || newCost < *stepData.LocDataMap[neighbor].Cost {
		hValue := stepData.Heuristic.Distance(stepData.GoalCoord, neighbor)
		breadcrumb := core.Crumb{To: neighbor, From: currentLoc.Breadcrumb}
		miniLoc := cq.MiniLoc{Breadcrumb: breadcrumb, Coord: neighbor, Cost: newCost}
		heap.Push(&stepData.Queue, cq.PrioBundle[cq.MiniLoc]{Priority: newCost + hValue, Item: miniLoc})
	}
}
