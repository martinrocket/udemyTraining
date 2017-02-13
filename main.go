package main

import (
	"fmt"
	//"time"
)

func main() {

	fmt.Println("Program Start....")
	makeCall()
}

func makeCall() {
	for i := 1; i < 1000; i ++ {
		fmt.Printf("%d \t %q \n", i, i)
	}
}
