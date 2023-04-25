package main

import (
	"fmt"
	"time"
)

func mySleeper(sec int) {
	dur := time.Duration(sec) * time.Second

	fmt.Println("Sleeping for", sec)
	select {
	case <-time.After(dur): // про этот канал я уже говорил. реализация простая.
		fmt.Println("Ending sleep")
		return
	}
}

func main() {

	sleeper := 2

	mySleeper(sleeper)
}