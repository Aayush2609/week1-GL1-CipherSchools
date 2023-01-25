package main

import "fmt"

/*
1. if-else statement
	if(condition){
		//statements
	}

	if(condition){
		//statements
	} else {
		//statements
	}

	if(condition){
		//statements
	}
	else if(condition){
		//statements
	}


2. switch(data){
	case var1:
		//statements
	case var2:
		//statements
	default:
		//statements
}

*/




func main() {
	// data := 10
	// fmt.Print("Enter a number")
	// var input int
	// fmt.Scanln(&input)
	
	// if input%2==0 {
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("you are odd")
	// }

	// if x:=10; x%2==0{
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("you are odd")
	// }

		data := 14567890
		switch data {
		case 10:
			data = 100
			fmt.Println(data)
		case 20:
			data = 103
			fmt.Println(data)
		case 11:
			data = 1002
			fmt.Println(data)
			fallthrough //execute the next case also
		case 15:
			data = 10000
			fmt.Println(data)
			fallthrough
		default:
			data = 10909
			fmt.Println(data)
		}




}