package Strategy

type Navigator struct {
	Strategy Strategy
}

func (n *Navigator) SetStrategy(s Strategy) {
	n.Strategy = s
}

func (n *Navigator) Route(startPoint, endPoint int) {
	n.Strategy.Route(startPoint, endPoint)
}
