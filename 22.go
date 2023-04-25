package main

import (
	"fmt"
	"os"
	"log"
	"math/big"
)

// Используем big для работы с большими числами (оптимальнее встроенных типов)

func divide(n1, n2 *big.Float) {
	fmt.Println(big.NewFloat(0).Quo(n1, n2))
}

func plus(n1, n2 *big.Float) {
	fmt.Println(big.NewFloat(0).Add(n1, n2))
}

func minus(n1, n2 *big.Float) {
	fmt.Println(big.NewFloat(0).Sub(n1, n2))
}

func multiply(n1, n2 *big.Float) {
	fmt.Println(big.NewFloat(0).Mul(n1, n2))
}

func main() {

	a := big.NewFloat(0) 
	b := big.NewFloat(0) 
	
	_, ok := a.SetString(os.Args[1])
	
	if !ok {
		log.Fatal("err a")
	}
	
	_, ok = b.SetString(os.Args[2])
	
	if !ok {
		log.Fatal("err b")
	}
	
	operation := os.Args[3]
	
	if operation == "" {
		log.Fatal("no op")
	}

	switch operation {
	case "div":
		divide(a, b)
	case "add":
		plus(a, b)
	case "minus":
		minus(a, b)
	case "multiply":
		multiply(a, b)
	default:
		fmt.Println("Такой функционал не поддерживается")
	}

}