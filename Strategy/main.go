package Strategy

var (
	start      = 10
	end        = 100
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&WalkStrategy{},
		&RoadStrategy{},
	}
)

func main() {
	nav := Navigator{}

	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}

}
