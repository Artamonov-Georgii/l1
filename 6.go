package main

import (
	"fmt"
	"time"
	"context"
)

func main() {

	// Первый способ - с помощью булевого флага

	chanOne := make(chan struct{}) // Этот канал сделан для того, чтобы программа демонстрировала один способ за раз.
	quitChanOne := make(chan bool)
	go func() {
		for {
			select {
			case <-quitChanOne:
				close(chanOne) // Закрываем канал, чтобы программа пошла дальше
				return
			default:
				// Каждую секунду - принтит
				fmt.Println("Мяу")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	quitChanOne <- true // Отсылаем булевый флажок в бесконечный фор-луп
	<-chanOne // А вот собственно и стоппер
	fmt.Println("\n================================")

	// Второй способ - с помощью закрытия канала, который определен принимать struct{} (потому что весит 0 байт)

	chanTwo := make(chan struct{})
	quitChanTwo := make(chan struct{})

	go func() {
		for {
			select {
			case <-quitChanTwo: // Закрытие канала отсылает zero value, struct{}{}
				close(chanTwo)
				return
			default:
				// Каждую секунду - принтит
				fmt.Println("Мяу")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	close(quitChanTwo) // Это работает потому, что из закрытого канала можно прочитать zero-value типа, в нашем случае пустую структуру
	<-chanTwo
	fmt.Println("\n================================")

	// С помощью контекста

	ctx, cancel := context.WithCancel(context.Background()) // context.Background - изначальный non-nil контекст. WithCancel дает функцию, которая отменяет контекст.
	chanThree := make(chan struct{})

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				close(chanThree)
				return
			default:
				// Каждую секунду - принтит
				fmt.Println("Мяу")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel() // Делает так, чтобы из ctx.Done() можно было прочитать значение
	<-chanThree
	fmt.Println("\n================================")

	// С помощью таймаута Time.After chan

	chanFour := make(chan struct{})
	randChan := make(chan int)

	go func() {
		for {
			select {
			case <-time.After(time.Second*10): // Если горутина не успевает за 10 секунд справится с задачей, или если у нее нет никаких новых данных 
				close(chanFour) // в течение n количества времени, то канал выдает значение, и горутину можно завершить
				return
			case i:=<-randChan:
				fmt.Println(i)
			}
		}
	}()
	
	randChan<-1
	randChan<-2
	randChan<-3

	<-chanFour
	fmt.Println("\n================================")

	// Есть еще пакет tomb, но с ним я не знаком
}