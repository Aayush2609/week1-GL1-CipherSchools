package main

import "fmt"

func main() {
	rec(4)
	// fmt.Println(fact(10))
	// fmt.Println(fib(10))
}

//Recursion basic question
func rec(num int) {
	if num <= 0 {
		return
	}
	// if num%2 == 0 {
		
	// 	rec(num - 1)
	// 	fmt.Println(num - 1)
	// } else {
		
	// 	rec(num - 1)
	// 	fmt.Println(num - 1)
	// }
	rec(num-1)
	rec(num-2)
	fmt.Println(num-1)
}


// Factorial of a number

func fact(number int) int {
	if number == 1 || number == 0{   // base case
		return 1
	}

	if number < 0 {				// corner case
		fmt.Println("Invalid Number")
		return -1
	}


	result := number * fact(number-1)   //here the function is calling itself within itself
	return result
}

// Fibonacci of a number

func fib(number int) int{
	if number == 0 || number == 1{      //base cases
		return number
	}
		result:=fib(number-1)+fib(number-2)
	return result
}


