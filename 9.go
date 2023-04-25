package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"os/signal"
	"time"
	"syscall"
)

// Я выполнял уже конвейер в предыдущих заданиях, поэтому здесь повторение

// Воркер, который читает из главного канала, умножает на два, и отправляет в канал для принтинга

func worker(readChan <-chan int, twoChan chan<- int, wg *sync.WaitGroup) {
	var data int
	var ok bool

	for {
		data, ok = <-readChan
		if !ok {
			fmt.Printf("Воркер закончил работу\n")
			close(twoChan)
			wg.Done()
			return
		}
		twoChan <- data*2
	}
}

// Принтер читает из канала и принтит результат умножения в stdout

func printer(twoChan <-chan int, wg *sync.WaitGroup) {
	var data int
	var ok bool

	for {
		data, ok = <-twoChan
		if !ok {
			fmt.Printf("Принтер закончил работу\n")
			wg.Done()
			return
		}
		fmt.Println(data)
	}
}

func main() {

	wg := &sync.WaitGroup{}
	
	readCh := make(chan int)
	twoCh := make(chan int)

	signalCh := make(chan os.Signal, 1) 
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	wg.Add(2)
	go worker(readCh, twoCh, wg)
	go printer(twoCh, wg)

	go func() {
		for {
			select {
			case <-signalCh:
				fmt.Println("Получен сигнал, завершаем работу")
				close(readCh)
				return
			case readCh <- rand.Intn(4):
				time.Sleep(time.Second)
			}
		}
	}()


	wg.Wait() // Схема graceful shutdown такая же
	fmt.Println("Все работники завершили работу, пока...")
}
