package main

import "fmt"

func findLeader(arr []int, size int) {
	max := arr[size-1]
	fmt.Println(max)
	for i := size-2; i >= 0; i-- {
		if arr[i] > max {
			max = arr[i]
			fmt.Println(arr[i])
		}
	}
}

func leader() {
	fmt.Println("LEADER")
	arr := []int{14, 12, 70, 15, 99, 65, 21, 90}
	findLeader(arr, len(arr))
}
