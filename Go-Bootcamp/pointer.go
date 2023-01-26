package main

import "fmt"

func main() {
	// i := 10
	// j := &i
	// *j = 100

	var i int
	i = 10
	// var j *int    //declaration.. j is empty

	j := new(int)   //declaration + assignment(j is not empty)
	j = &i
	*j = 100

	// fmt.Println(j)
	// fmt.Println(i)

	name := new(string)
	// name = "abc"
	fmt.Println(name)
}

