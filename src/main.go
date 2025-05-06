package main 

import (
	"fmt"
)




func main (){
		// asking for a weight 
	fmt.Println("Introduce the weight to be prepared")

	fmt.Scanf("%v \n", &weight)

	fmt.Printf("El peso es: %v\n", weight)
	CalcDisks(weight)

}
