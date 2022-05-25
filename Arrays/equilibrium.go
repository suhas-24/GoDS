package main

import "fmt"

func sum(arr []int)int{
	sum:=0
	for _,num := range arr{
		sum+= num
	}
	return sum
}
func findOdd(arr []int)int{
	totalSum := sum(arr)
	sumSoFar := 0
	for i := range arr{
		if sumSoFar == totalSum - sumSoFar{
			return i
		}
	sumSoFar += arr[i]
	}
	return -1
}
func equilibrium(){
	fmt.Println("-------------")
	fmt.Println("Equilibrium")
	arr := []int{1,4,5}
	i := findOdd(arr)
	if i!= -1{
		fmt.Println(arr[:i])
		fmt.Println(arr[i:])
	}else{
		fmt.Println("Cannot be partitioned")
	}
}