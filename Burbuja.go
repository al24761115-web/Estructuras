package main

import "fmt"

func bubbleSort(arr []int, ascending bool) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ascending {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			} else {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	return arr
}

func main() {
	original := []int{64, 34, 25, 12, 22, 11, 90}

	asc := make([]int, len(original))
	copy(asc, original)
	bubbleSort(asc, true)
	fmt.Println("Ascendente:", asc)

	desc := make([]int, len(original))
	copy(desc, original)
	bubbleSort(desc, false)
	fmt.Println("Descendente:", desc)
}
