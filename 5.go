package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Использую такой же темплейт, как в предыдущем задании

func worker(readChan <-chan int, workerNum int, wg *sync.WaitGroup) {
	var data int
	var ok bool

	for {
		data, ok = <-readChan
		if !ok {
			fmt.Printf("Воркер %d закончил работу\n", workerNum)
			wg.Done()
			return
		}
		fmt.Println(data)
	}
}

func main() {
	
	if len(os.Args)<2 { // Теперь из командной строки будет парсится количество секунд N
		log.Fatal("Нет количества секунд")
	}

	numWorkers := 5 // Пусть будет 5 воркеров. Случайное число.
	seconds, err := strconv.Atoi(os.Args[1])
	
	if err != nil || seconds == 0 {
		log.Fatal("Неправильное количество секунд") // Ловим ошибку
	}

	mainCh := make(chan int)
	signalCh := make(chan int) // Тут уже не принципиально с буфером

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(mainCh, i, &wg)
	}
	
	duration := time.Second * time.Duration(seconds)

	go func() {
		for {
			select {
			case <-signalCh:
				fmt.Println("Получен сигнал, завершаем работу")
				close(mainCh)
				return
			case mainCh <- rand.Int():
				time.Sleep(time.Second)
			}
		}
	}()

	fmt.Println("\n \n \n Начинаем отсчет \n \n \n") // Сон отсюда
	time.Sleep(duration)	
	close(signalCh) // Механизм завершения такой же

	wg.Wait()
	fmt.Println("Все работники завершили работу, пока...")
}
