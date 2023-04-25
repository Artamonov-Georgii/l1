package main 

import (
	"fmt"
)

func deleteElement(slice []int, index int) []int {
    copy(slice[index:], slice[index+1:]) // копируем в slice[index:] все, кроме самого элемента index
    return slice[:len(slice)-1]
}

func main() {
    aSlice := []int{1, 2, 3, 4}
    aSlice = deleteElement(aSlice, 2) // Важно возвращать! Иначе потеряем все... (изменения не передаются во вне)
    fmt.Println(aSlice)
}
