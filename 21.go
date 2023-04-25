package main

import (
	"fmt"
)

// Адаптер - это когда мы хотим реализовать метод, который не определен для структуры,
// путем создания отдельной структуры адаптера, куда адаптируемая структура передается в виде параметра,
// и определения этого метода уже для адаптера
// Здесь немного игрушечный пример. Обычно адаптеры намного сложнее

type Laugher interface {
	Laugh()
}

type LaughAdaptee struct {
	name string
	age  int
}

func (a *LaughAdaptee) Giggle() {
	a.age--
}

type LaughAdapter struct {
	adaptee *LaughAdaptee	
}

func (la *LaughAdapter) Laugh() {
	la.adaptee.Giggle()
}

func main() {
	adaptee := &LaughAdaptee{"Vova", 42}
	adapter := &LaughAdapter{adaptee}

	adapter.Laugh()

	fmt.Println(adaptee)

}