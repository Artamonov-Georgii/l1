package main

import "fmt"


// Простой метод через type assertion
// interface{} имплицитно выполняется любым объектом в го. С помощью reflection он дает понять, с какой переменной имеем дело

func getType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	case bool:
		fmt.Println("Type: bool")
	case chan int:
		fmt.Println("Type: channel")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
    var i interface{}

    i = 42
    getType(i)

    i = "hello"
    getType(i)

    i = true
    getType(i)

    ch := make(chan int)
    i = ch
    getType(i)
}

