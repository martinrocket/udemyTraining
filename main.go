package main

import (
	"fmt"
	//"time"
)

func main() {
	myInt := 10
	fmt.Println("Program Start....")
	makeCall()
	fmt.Printf("%d \t %v \n", &myInt, &myInt)
	fmt.Println(&myInt)
}

func makeCall() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d \t %q \n", i, i)
	}

}
