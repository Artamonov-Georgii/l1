package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Counter struct {
	Count int
	m sync.Mutex // Чтобы не было lost update или других проблем, связанных с конкурентностью
}

func newCounter() *Counter {
	return &Counter{} // Конструктор
} 

func (c *Counter) incrementCounter(by int) {
	c.m.Lock()
	c.Count++
	c.m.Unlock()
}

func main() {

	counter := newCounter()

	for i:=0; i<10; i++ {
		go func() {
			for {
				counter.incrementCounter(rand.Intn(10))
				time.Sleep(time.Second)
				fmt.Println(counter.Count)
			}
		}()
	}


	time.Sleep(time.Second*20)
	fmt.Println(counter.Count)
}