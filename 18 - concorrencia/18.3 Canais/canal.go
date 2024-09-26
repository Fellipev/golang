package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Hello", canal)

	for {
		msg := <-canal
		fmt.Println(msg)
	}

}

func escrever(txt string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- txt
		time.Sleep(time.Second)
	}
}
