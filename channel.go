package main

import (
	"fmt";
	"time";
)

func ping(c chan string)  {
	for index := 0; ; index++ {
		c <- fmt.Sprintf("Ping %d", index)
	}
}

// OR

func pingProtected(c chan<- string)  {
	for index := 0; ; index++ {
		c <- fmt.Sprintf("Ping protected %d", index)
	}
}

func print(c chan string)  {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second + 1)
	}
}

// OR
func printProtected(c <-chan string)  {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second + 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go print(c)

	// go pingProtected(c)
	go printProtected(c)

	fmt.Scanln()
}
