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
			 rest_weight -= (2.00*Big_red)
			fmt.Println("DISCOS ROJOS GRANDES")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		// disco azul grande
		}else if (rest_weight - (2.00*Big_blue)) >= 0 {
			 rest_weight -= (2.00*Big_blue)
			fmt.Println("DISCOS AZULES GRANDES")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		// disco amarillo grande
		}else if (rest_weight - (2.00*Big_yellow)) >= 0 {
			 rest_weight -= (2.00*Big_yellow)
			fmt.Println("DISCOS AMARILLOS GRANDES")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		// disco verde grande
		}else if (rest_weight - (2.00*Big_green)) >= 0 {
			 rest_weight -= (2.00*Big_green)
			fmt.Println("DISCOS VERDES GRANDES")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================") 
		}else if (rest_weight - (2.00*Big_white)) >= 0 {
			 rest_weight -= (2.00*Big_white)
			fmt.Println("Discos blancos GRANDES")
			fmt.Printf("QUEDAN %v\n", rest_weight)
			fmt.Println("=====================================")
		// discos peque rojo
		}else if (rest_weight - (2.00*Small_red)) >= 0 {
			 rest_weight -= (2.00*Small_red)
			fmt.Println("DISCOS ROJOS PEQUE")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		// discos azules peque
		}else if (rest_weight - (2.00*Small_blue)) >= 0 {
			 rest_weight -= (2.00*Small_blue)
			fmt.Println("DISCOS AZUL PEQUE")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		// discos amarillos
		}else if (rest_weight - (2.00*Small_yellow)) >= 0 {
			 rest_weight -= (2.00*Small_yellow)
			fmt.Println("DISCOS AMARILLOS PEQUE")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		// discos verdes
		}else if (rest_weight - (2.00*Small_green)) >= 0 {
			 rest_weight -= (2.00*Small_green)
			fmt.Println("DISCOS VERDES PEQUE")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		// discos blancos peques
		}else if (rest_weight - (2.00*Small_white)) >= 0 {
			 rest_weight -= (2.00*Small_white)
			fmt.Println("DISCOS BLANCOS PEQUE")
			fmt.Printf("Peso Restante: %v\n", rest_weight)
			fmt.Println("=====================================")
		}else{
			break
		}
	}
}
