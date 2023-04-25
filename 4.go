package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

// Создадим функцию воркера
// Воркер читает данные из канала, имеет свой собственный номер и доступ к вейтгруппе (по поинтеру конечно, чтобы не копировалась переменная)
// Можно было конечно и структуру объявить, но это уже чересчур

func worker(readChan <-chan int, workerNum int, wg *sync.WaitGroup) {
	var data int
	var ok bool

	for {
		data, ok = <-readChan // Он проверяет каждую итерацию, закрыт ли канал. Если канал закрыт, то return
		if !ok {
			fmt.Printf("Воркер %d закончил работу\n", workerNum)
			wg.Done() // Отнимаем единицу из wg
			return
		}
		fmt.Println(data)
	}
}

func main() {
	if len(os.Args)<2 {
		log.Fatal("Нет количества воркеров") // Парсит аргументы из командной строки. Первый аргумент - кол-во воркеров.
	}

	numWorkers, err := strconv.Atoi(os.Args[1]) // Ascii to integer из значения командной строки
	
	if err != nil || numWorkers == 0 {
		log.Fatal("Неправильное количество воркеров")
	}

	mainCh := make(chan int) 
	signalCh := make(chan os.Signal, 1) // Важно сделать канал с буфером, чтобы не терялись сигналы (референс - https://www.reddit.com/r/golang/comments/onwoi3/channels_of_size_1/)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM) 

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(mainCh, i, &wg)
	}

	go func() {
		for {
			select {
			case <-signalCh: // Здесь важно иметь буфер, так как мы можем нажать на ctrl c, когда канал не готов к чтению
				fmt.Println("Получен сигнал, завершаем работу")
				close(mainCh) // Отсюда начинается graceful shutdown. 
				return // Сначала закрывается поток - потом, !!после завершения текущей итерации!!, закрываются воркеры - потом wg.Done() и пока
			case mainCh <- rand.Int():
			}
		}
	}()

	wg.Wait() // Ждем, пока все воркеры перестанут работать...
	fmt.Println("Все воркеры завершили работу, пока...")
}
