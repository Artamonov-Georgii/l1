package main // Этот код не будет использовать sync.Map. Предположим, что мы хотим блокировать всю мапу перед записью

import (
	"fmt"
    "math/rand"
    "sync"
    "time"
)

type safeMap struct {
	theMap   map[int]int
	mutex    sync.RWMutex // Запихиваем сюда RWMutex, чтобы можно было читать, но не писать.
}

func newMap() *safeMap {
	return &safeMap{theMap: map[int]int{}} // Создаем обычную функцию конструктор 
}

func (m *safeMap) setKey(key, value int) { // Делаем функцию, устанавливающую значение
	m.mutex.Lock() // Делает write block, чтобы не было проблем
	m.theMap[key] = value
	m.mutex.Unlock()
}

func main() {

    concurentMap := newMap()
    signalCh := make(chan int)
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int, wg *sync.WaitGroup) {
            for {
                select {
                case <-signalCh:
                    wg.Done()
                    return
                default:
					randValue := rand.Intn(10) // Здесь должна происходить трудозатратная операция
                    concurentMap.setKey(randValue, randValue*2+i) // А вот здесь уже легкая, чтобы блокирование не было супер дорогим
                } // Добавляю i чтобы продемонстрировать, что работают разные горутины
            }
        }(i, &wg)
    }

    duration := time.Second * time.Duration(5)
    time.Sleep(duration)

    for i := 0; i < 5; i++ {
        signalCh <- 1 // Ради разнообразия закончим горутины таким образом
    }

	fmt.Println(concurentMap.theMap)
    wg.Wait()
}
