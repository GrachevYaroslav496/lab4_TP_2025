package Strategy

import "fmt"

type RoadStrategy struct {
}

func (s *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 60
	trafficJam := 4
	totalRoad := endPoint - startPoint
	timeInRoad := 30
	totalTime := totalRoad * trafficJam * timeInRoad
	fmt.Printf("Road A: [%d] to B: [%d]\nAvg speed: [%d] Traffic jam: [%d] TotalRoad: [%d] Total time: [%d] Time in road: [%d]\n", startPoint,
		endPoint, avgSpeed, trafficJam, totalRoad, totalTime, timeInRoad)
}
