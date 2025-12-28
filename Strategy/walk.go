package Strategy

import "fmt"

type WalkStrategy struct {
}

func (s *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4
	totalRoad := endPoint - startPoint
	timeInRoad := 75
	totalTime := totalRoad * timeInRoad
	fmt.Printf("Road A: [%d] to B: [%d]\nAvg speed: [%d] TotalRoad: [%d] Total time: [%d]\n\n", startPoint,
		endPoint, avgSpeed, totalRoad, totalTime)
}
