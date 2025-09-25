package main

import "fmt"

func main() {
	number, eur, usd := output()
	result := countCurrencies(number, eur, usd)
}
func output() (float64, float64, float64) {
	var numberCurrency, EUR, USD float64
	fmt.Scan(&numberCurrency, &EUR, &USD)
	return numberCurrency, EUR, USD
}
func countCurrencies(numbers, EUR, USD float64) float64 {

}
