package main

import (
	"fmt"
	"time"
)

func say(words string) {
	for i := 0; i < 5; i++ {
		fmt.Println(words)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go say("go routine")
	go say("not a go routine")
	fmt.Scanln()
}
