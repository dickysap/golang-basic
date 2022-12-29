package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("============ Multiple Value ================")
fmt.Println(MultipleFunction(2))
fmt.Println("============ Return Value ================")
fmt.Println(ReturnValue(2))
fmt.Println("============ Variadic ================")
fmt.Println(VariadicFunc(2,2,2,2))

// Closure Function
var getMinMax = func(n []int)(int, int){
	var min , max int 

	for i,value := range n {
		switch {
		case i == 0 :
			max, min = value, value
		case value > max :
			max = value
		case value < min :
			min = value
		}
	}
	return min, max
}
var numbbers = []int{2,3,4,5,6,7,8,2,1,3,2,1}
var min, max = getMinMax(numbbers)
fmt.Printf("data	:	%v\nmin	:	%v\nmax	:	%v\n", numbbers, min, max)

fmt.Println("============ Expresion Function ================")
ExpressionFunc()
}

func MultipleFunction(d float64) (float64, float64) {
	var pow = math.Pow(d / 2, 2)
	fmt.Println(pow)
	var area = math.Phi * pow
	var keliling = math.Phi * d

	return area, keliling

}

func ReturnValue(d float64) (area float64, keliling float64){
	var pow = math.Pow(d / 2, 2)
	fmt.Println(pow)
	area = math.Phi * pow
	keliling = math.Phi * d

	return
}

//variadic function
func VariadicFunc(numbers ...int)int {
								// hasil return bertipe Int 

	var total int = 0
	for _, value := range numbers {
		total += value
	}

	return total
}


//Immediately-Invoked Function Expression (IIFE)

func ExpressionFunc(){
	var numbers  = []int{2,3,4,2,1,1,3,4,3,2,5,6,3,2}

	var newNumbers = func(min int)[]int{
		var r []int
		
		for _, value := range numbers {
			if value < min {
				continue
			}
			r = append(r, value)
		}
		return r
	}(3)

	fmt.Println("Original Number	:", numbers)	
	fmt.Println("Filtered Number	:", newNumbers)	
}