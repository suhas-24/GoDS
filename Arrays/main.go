package main

import "fmt"

func main() {
	createArray()
}

func createArray() {
	var arr [5]int // declares a variable arr as an array of five integers
	arr[0] = 10 
	var arr2 [3]string = [3]string{"Hello", "World", "!"} // partial assignment using array literals
	arr3 := [...]float64{0.12, 2.24, 3.36 } // initializing an array using ellipses...
	var arr4 = make([]bool, 5) // using make() function
	arr4[0] = true 
	fmt.Println("Array of integers : ", arr)
	fmt.Println("Array of strings : ", arr2)
	fmt.Println("Array of floating values : ", arr3)
	fmt.Println("Array of boolean values : ", arr4)
}
