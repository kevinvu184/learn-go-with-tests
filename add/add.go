package main

import "log"

func main() {
	log.Println(add(1, 2, 3))
}

func add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
