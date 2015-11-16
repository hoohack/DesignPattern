package main

import "fmt"
import "./travel"

func main() {
	strategyAir := new(travel.AirplaneStrategy)
	person := travel.SetStrategy(strategyAir)
	fmt.Printf(person.Travel())

	strategyTra := new(travel.TrainStrategy)
	person = travel.SetStrategy(strategyTra)
	fmt.Printf(person.Travel())

	strategyBi := new(travel.BicycleStrategy)
	person = travel.SetStrategy(strategyBi)
	fmt.Printf(person.Travel())
}
