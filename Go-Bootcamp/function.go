package main

import "fmt"

/*
func function_name(){
	statement -1
	statement -2
	...
}


*/

//func is used to create the function
func main(){
result, x := c(10, "go")
fmt.Println(result)
fmt.Println(x)
r, _ := b(1,2,4,5,6,7,8)
fmt.Println(r)

}

func a() (int, string){
	return 122, "wedf"
}
func b(args ...int)(bool, int){  //... before the datatype defines that multiple arguments will be taken in that declared variable
	for _, v := range args {
		fmt.Print(v)
	}
	return true, 11
}
func c (w int, name string) (i int, j string){
	i = 10
	j = "rahul"
	return
}
	
