package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

func main() {
	arr := []int{9, 2, 10, 12, 3, 7, 8, 4, 6, 5, 1}
	fmt.Println("Initial array is:", arr)
	sortedArr := QuickSort(arr)
	fmt.Println("Sorted array is:", sortedArr)
}
