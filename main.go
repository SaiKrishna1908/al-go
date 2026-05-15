package main

import (
	"fmt"

	"github.com/SaiKrishna1908/al-go/internal/algo/search"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3

	target_pos := search.BinarySearch(arr, target)

	fmt.Println(target_pos)
}
