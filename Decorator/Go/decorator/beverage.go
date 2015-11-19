package decorator

type beverage interface {
	Cost() float64
	GetDescription() string
}
