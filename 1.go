package main

import (
	"fmt"
)

// Пусть есть структура Human

type Human struct {
    name    string
    surname string
    age     int
    // Здесь могут быть определены другие параметры
}

// А для нее определен метод SayHello, который печатает параметры инстанса структуры

func (h Human) SayHello() {
    fmt.Printf("Имя: %s %s \nВозраст: %d лет\n", h.name, h.surname, h.age)
}

// Пусть есть структура Action

type Action struct {
    Human
    ActionType string
    // Здесь могут быть определены другие структуры или другие параметры
}

// А для нее определен метод DoAction, который внутри вызывает метод SayHello

func (a Action) DoAction() {
    fmt.Printf("DoAction call\n")
    a.SayHello() // SayHello не определен для Action, поэтому по дефолту будет воспроизводится метод из структуры Human.
    fmt.Printf("SayHello call\n") // Если бы для Action был определен такой же метод, 
                                  // то надо было бы писать a.Human.SayHello(), чтобы вызвать метод из субструктуры
}

func main() {
    human := Human{name: "Вова", surname: "Вовов", age: 18}
    action := Action{human, "Action 1"}
    action.DoAction()
}