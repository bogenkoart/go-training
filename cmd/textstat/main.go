package main

import (
	"fmt"

	"github.com/bogenkoart/go-training/internal/slicework"
)

func main() {
	fmt.Println(slicework.Dedup([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(slicework.Chunk([]int{1, 2, 3, 4}, 2))
	fmt.Println(map[string]int{"work": 1, "work1": 2}) 
}
