package main

import "fmt"

func main() {
	// number := 10
	var number int
	fmt.Println(number)



	//store a function as a value to a variable
	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In a function")
		return 10 + x
	}
	getInt(1)
	// getInt = func(x int) int {
	// 	fmt.Println("In another function")
	// 	return 20 + x
	// }


	var y func()
	y = func() {
		fmt.Println(y)
	}

	j := func(x int)int {  //declaring and calling a function at the same place
		fmt.Println("In another function")
		return 20 + x   
	}(19)
	fmt.Println(j)

}
func g(getInt func(int) int){   //calling a function into another function and using it a s a parameter when needed
	getInt(78)
}
//function is treated as the first class member in golang