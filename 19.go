package main 

import (
	"fmt"
	"strings"
)

// Буду использовать стрингбилдер, самый эффективный способ конкатенации строк)

func stringReverser(input string) string {

	if len(input) == 0 || len(input) == 1 { 
		return input
	}

	var builder strings.Builder

	i := len(input)-1

	for i >= 0 { // Просто задом наперед по буквам идем
		builder.WriteRune(rune(input[i]))
		i--
	}

	result := builder.String()
	return result
}

func main() {
	stringToReverse := "reverseMe"
	
	fmt.Println(stringReverser(stringToReverse))
}