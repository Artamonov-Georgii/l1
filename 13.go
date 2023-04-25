package main

import (
	"fmt"
)

// Странный метод... Первое, что пришло в голову

func swapWithoutTemp(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	a := 1
	b := 2

	swapWithoutTemp(&a, &b)
	fmt.Println(a, b)

	// a, b := b, a создает временную tuple, поэтому хз, считается ли
}