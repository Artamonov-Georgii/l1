package main

import (
	"fmt"
	"sync"
)


func main() {
	// Можно было сделать и слайс, но раз уж указано массив, а не динамический массив...
    arr := [5]int{2, 4, 6, 8, 10}
    sqArr := [5]int{}
    var wg sync.WaitGroup // чтобы не завершалась преждевременно программа добавляем вейтгруппу

    for i := 0; i < len(arr); i++ {
        wg.Add(1)
        go func(i int) { // важно подавать i в качестве параметра, иначе он будет принимать только последнее значение i
            sqArr[i] = arr[i] * arr[i]
            wg.Done()
        }(i)
    }

    wg.Wait()
    fmt.Println(sqArr)
}
