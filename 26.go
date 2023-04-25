package main

import (
	"fmt"
	"unicode"
)

func uniquenessCheck(s string) bool {
	runeMap := make(map[rune]struct{})

	for _, v := range s {
		_, ok := runeMap[unicode.ToLower(rune(v))] // аналог .lower() python
		if ok {
			return false // Хоть один раз увидем одно и тоже - уже нет смысла дальше что то выполнять, возвращаем фолс
		}
		runeMap[unicode.ToLower(rune(v))] = struct{}{}
	}

	return true
}

func main() {
	stringerChecker := "chestringC"

	fmt.Println(uniquenessCheck(stringerChecker))
}