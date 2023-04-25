package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	resultSum := 0 // Сюда будем прибавлять результат
	var wg sync.WaitGroup
	var mutex sync.Mutex // Создаем мьютекс, дабы избежать потенциально lost update проблему

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(i int) {
			g := arr[i] * arr[i]
			mutex.Lock() // Блокируем мьютекс перед доступом к resultSum
			resultSum += g
			mutex.Unlock() // Разблокируем мьютекс после доступа к resultSum
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(resultSum)
}
