package main

import "fmt"

func main() {
	const USDonEUR = 0.85
	const USDonRUB = 83.49
	EURonRUB := (1 - USDonEUR) + USDonRUB
	fmt.Print("Евро в рублях =", EURonRUB)

}
