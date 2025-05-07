package main 

import (
	"fmt"
)




func main (){
	
	// asking for a weight 
	fmt.Println("Introduce the weight to be prepared")

	fmt.Scanf("%v \n", &weight)
	fmt.Printf("El peso es: %v\n", weight)

	// indicate if is Male o Female contest
	fmt.Println("Which Contest it is: Type (F) Female, (M) if Male")
	fmt.Scanf("%v \n", &Sex_category)		
	
	// sending the bar with the correct category
	
	if Sex_category == "F" {
		CalcDisks(weight, Women_bar)
	}else if Sex_category == "M" {
		CalcDisks(weight, Men_bar)
	}else {
		err := "The category is not allowed. Try again"
		panic(err)
	}
}
