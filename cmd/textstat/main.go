package main

import (
	"fmt"
	"github.com/bogenkoart/go-training/internal/textstat"
)

func main() {
	var word string
	fmt.Scan(&word)
	fmt.Println(textstat.Frequency(textstat.Words(word)))
}
