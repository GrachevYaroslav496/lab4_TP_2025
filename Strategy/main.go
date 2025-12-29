package main

import (
	"OtherCode/TP/lab4_TP_2025/Strategy"
)

var (
	start      = 10
	end        = 100
	strategies = []Strategy.Strategy{
		&Strategy.PublicTransportStrategy{},
		&Strategy.WalkStrategy{},
		&Strategy.RoadStrategy{},
	}
)

func main() {
	nav := Strategy.Navigator{}

	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}

}
