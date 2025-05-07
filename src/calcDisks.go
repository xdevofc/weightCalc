package main 

import (
	"fmt"
)

func CalcDisks (weight float64, bar float64) {

//	var flag bool = true

	// start calculating

	// resting the reglamentary implements
	rest_weight := weight - bar - (2*Collars)

	fmt.Printf("EL PESO ES: %v\n", rest_weight)
	
	for rest_weight != 0 {
		// disco rojo grande
		if (rest_weight - (2.00*Big_red)) >= 0 {
			rest_weight = rest_weight - (2.00*Big_red)
			fmt.Println("DISCOS ROJOS GRANDES\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		// disco azul grande
		}else if (rest_weight - (2.00*Big_blue)) >= 0 {
			rest_weight = rest_weight - (2.00*Big_blue)
			fmt.Println("DISCOS AZULES GRANDES\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		// disco amarillo grande
		}else if (rest_weight - (2.00*Big_yellow)) >= 0 {
			rest_weight = rest_weight - (2.00*Big_yellow)
			fmt.Println("DISCOS AMARILLOS GRANDES\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		// disco verde grande
		}else if (rest_weight - (2.00*Big_green)) >= 0 {
			rest_weight = rest_weight - (2.00*Big_green)
			fmt.Println("DISCOS VERDES GRANDES\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
	 
		}else if (rest_weight - (2.00*Big_white)) >= 0 {
			rest_weight = rest_weight - (2.00*Big_white)
			fmt.Println("Discos blancos GRANDES\n")
			fmt.Printf("QUEDAN %v\n", rest_weight)
		// discos peque rojo
		}else if (rest_weight - (2.00*Small_red)) >= 0 {
			rest_weight = rest_weight - (2.00*Small_red)
			fmt.Printf("DISCOS ROJOS PEQUE\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		// discos azules peque
		}else if (rest_weight - (2.00*Small_blue)) >= 0 {
			rest_weight = rest_weight - (2.00*Small_blue)
			fmt.Printf("DISCOS AZUL PEQUE\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		// discos amarillos
		}else if (rest_weight - (2.00*Small_yellow)) >= 0 {
			rest_weight = rest_weight - (2.00*Small_yellow)
			fmt.Printf("DISCOS AMARILLOS PEQUE\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		// discos verdes
		}else if (rest_weight - (2.00*Small_green)) >= 0 {
			rest_weight = rest_weight - (2.00*Small_green)
			fmt.Printf("DISCOS VERDES PEQUE\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		// discos blancos peques
		}else if (rest_weight - (2.00*Small_white)) >= 0 {
			rest_weight = rest_weight - (2.00*Small_white)
			fmt.Printf("DISCOS BLANCOS PEQUE\n")
			fmt.Printf("QUEDAN: %v\n", rest_weight)
		}else{
			break
		}
	}
}
