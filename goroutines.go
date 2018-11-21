package main

import "fmt"

func routine(max int)  {
	fmt.Println("Salut")
	for index := 0; index < max; index++ {
		fmt.Println(index)
	}
}

func main()  {
	go routine(10)
	go routine(10)
	fmt.Scanln()
	fmt.Println("done")
}
