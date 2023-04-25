package main

import (
	"fmt"
)

func main() {
	set1 := map[string]bool{"apple": true, "orange": true, "banana": true}
	set2 := map[string]bool{"orange": true, "grape": true, "kiwi": true}

	intersection := make(map[string]struct{}) // Мапа нужна для того, чтобы лукап по ключю был константный

	for k := range set1 {
		if set2[k] { // Если такое значение есть, то помещаем его в мапу. 
			intersection[k] = struct{}{}
		}
	}

	fmt.Println(intersection)
}