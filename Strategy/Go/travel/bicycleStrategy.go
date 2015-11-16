package travel

type BicycleStrategy struct {
}

func (bs BicycleStrategy) traveler() string {
	return "Go to travel by bicycle\n"
}
