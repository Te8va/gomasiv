package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		numbers [5]int
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}
	fmt.Println("Начальный массив: ", numbers)
	sort := bubble_sort(numbers)
	fmt.Println("Конечный массив: ", sort)

}

func bubble_sort(array [5]int) [5]int {
	max := 0
	for i := 0; i < len(array)-1; i++ {
		for j := 1; j < len(array)-i; j++ {
			if array[j] < array[j-1] {
				max = array[j-1]
				array[j-1] = array[j]
				array[j] = max
			}
		}
	}
	return array
}
