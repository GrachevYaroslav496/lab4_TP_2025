package Strategy

import "fmt"

type PublicTransportStrategy struct {
}

func (pts *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 60
	totalRoad := endPoint - startPoint
	timeInRoad := 40
	totalTime := totalRoad * timeInRoad
	fmt.Printf("Road A: [%d] to B: [%d]\nAvg speed: [%d] TotalRoad: [%d] Total time: [%d]\n\n", startPoint,
		endPoint, avgSpeed, totalRoad, totalTime)
}
