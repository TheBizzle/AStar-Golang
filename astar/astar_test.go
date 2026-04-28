package astar_test

import (
	"fmt"
	"testing"

	astar "github.com/TheBizzle/AStar-Golang/astar"
	heur "github.com/TheBizzle/AStar-Golang/heuristic"
	sd "github.com/TheBizzle/AStar-Golang/stepdata"
	core "github.com/TheBizzle/PathFindingCore-Golang/pathingmap"
	testset "github.com/TheBizzle/PathFindingCore-Golang/testset"
)

func TestAStarAllMaps(t *testing.T) {
	heuristics := []heur.Heuristic{heur.Euclidean, heur.Manhattan, heur.Dijkstra}

	for i, testMap := range testset.Tests {
		for _, heur := range heuristics {
			t.Run(fmt.Sprintf("map%02d/%v", i+1, heur), func(t *testing.T) {
				result, stepData := astar.Run(testMap.MapStr, heur)

				isSolveable := testMap.Dist != nil
				if isSolveable {
					if result == core.SuccessfulRun {
						trail := stepData.LocDataMap[stepData.GoalCoord].Breadcrumb.Array()
						actual := len(trail) - 1
						expected := *testMap.Dist
						if float64(actual) != expected {
							t.Errorf("Wrong solution length.  Expected: %g, but got: %d.\n%s",
								expected, actual, formatFailureInfo(stepData))
						}
					} else {
						t.Errorf("Failed to solve map\n%s", formatFailureInfo(stepData))
						return
					}
				} else if result == core.SuccessfulRun {
					t.Errorf("Solved an unsolveable map")
				}
			})
		}
	}
}

func formatFailureInfo(stepData *sd.StepData) string {
	var breadcrumbStr string
	locData, ok := stepData.LocDataMap[stepData.CurrentCoord]
	if ok {
		breadcrumbStr = fmt.Sprintf("%v", locData.Breadcrumb)
	} else {
		breadcrumbStr = fmt.Sprintf("(no location data for CurrentCoord %v)", stepData.CurrentCoord)
	}

	return fmt.Sprintf(
		"iters so far: %d / %d  |  current coord: %v  |  breadcrumb: %v\nmap:\n%v",
		stepData.ItersSoFar,
		stepData.MaxIters,
		stepData.CurrentCoord,
		breadcrumbStr,
		stepData.PathingMap,
	)
}
