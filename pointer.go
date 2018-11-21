package main

import "fmt"

func setValue(ptr *int)  {
	*ptr = 100
}

func main() {
	x := 10
	setValue(&x)
	fmt.Println(x)
}
