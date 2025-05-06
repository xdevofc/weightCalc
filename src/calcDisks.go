package main 

import (
	"fmt"
)

func CalcDisks (weight float64) float64 {

	// start calculating
	fmt.Printf("El resultado desde calc es: %v \n", weight)

	rest_wieght := weight - Men_bar - Collars
	res := rest_wieght
	
	fmt.Printf("El resultado desde calc es: %v \n", res)


	return res
}
