package Strategy

type Navigator struct {
	strategy Strategy
}

func (n *Navigator) SetStrategy(s Strategy) {
	n.strategy = s
}

func (n *Navigator) Route(startPoint, endPoint int) {
	n.strategy.Route(startPoint, endPoint)
}
