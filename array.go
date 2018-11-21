package main

import "fmt"

func main()  {
	tab := [4]int{
		0,
		1,
		20,
		30,
	}

	for index := 0; index < len(tab); index++ {
		fmt.Println("Valeur", float64(tab[index]))
	}

	var total int
	for _, value := range tab {
		total += value
	}
}
