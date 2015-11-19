package main

import "fmt"
import "./decorator"

func main() {
	dark_roast := decorator.NewDarkRoast()
	espresso := decorator.NewEspresso()
	decaf := decorator.NewDecaf()
	house_blend := decorator.NewHouseBlend()

	mocha := decorator.NewMocha(dark_roast);
	double_mocha := decorator.NewMocha(mocha);

	espresso_soy := decorator.NewSoy(espresso)
	espresso_soy_mocha := decorator.NewMocha(espresso_soy)

	fmt.Printf("description\n%s === price: %.2f$\n\n", decaf.GetDescription(), decaf.Cost());
	fmt.Printf("description\n%s === price: %.2f$\n\n", house_blend.GetDescription(), house_blend.Cost());
	fmt.Printf("description\n%s === price: %.2f$\n\n", mocha.GetDescription(), mocha.Cost());
	fmt.Printf("description\n%s === price: %.2f$\n\n", double_mocha.GetDescription(), double_mocha.Cost());
	fmt.Printf("description\n%s === price: %.2f$\n\n", espresso_soy.GetDescription(), espresso_soy.Cost());
	fmt.Printf("description\n%s === price: %.2f$\n\n", espresso_soy_mocha.GetDescription(), espresso_soy_mocha.Cost());
}
