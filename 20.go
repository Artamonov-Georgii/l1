package main 

import (
	"fmt"
	"strings"
)

// Тоже реализуем через стрингбилдер

func reverseOrder(input string) string {

	var builder strings.Builder

	words := strings.Fields(input) // Эта функция сплитит по пробелу, аналог .split(" ") python
	
	//
	for i:=len(words)-1; i>=0; i-- {
		builder.WriteString(words[i])
		builder.WriteString(" ")
	}

	return builder.String()
}

func main() {
	inputStr := "hello world ok how are you"

	fmt.Println(reverseOrder(inputStr))
}