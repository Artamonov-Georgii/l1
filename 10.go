package main 

import (
	"fmt"
	"math"
)

// Вот два способа решения



func fetcher1(arr []float64) map[int][]float64 {

	resMap := make(map[int][]float64, len(arr)) // делаем капасити = длине слайса
	var key int
	
	for _, v := range arr {
		if v < 0 {
			key = int(math.Ceil(v / 10)) // если меньше нуля, то ключ - потолок результата деления значения на 10
		} else {
			key = int(math.Floor(v / 10)) // если больше нуля, то наоборот
		}
		fmt.Println(key, v)
		resMap[key*10] = append(resMap[key*10], v)
	}

	return resMap
}

func fetcher2(arr []float64) map[int][]float64 {

	resMap := make(map[int][]float64, len(arr)) 
	var key int

	for _, v := range arr {
		if v < 0 {
			key = -(int(-v+9)/10) // Это мы делаем из-за того, что округление идет вниз. Если хотим вверх, соответственно, надо прибавить делитель-1 (чтобы не было превышения) 
		} else {
			key = int(v) / 10
		}
		resMap[key*10] = append(resMap[key*10], v)
	}

	return resMap
}


func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res1 := fetcher1(array)
	res2:= fetcher2(array)

	fmt.Println(res1)
	fmt.Println(res2)
}