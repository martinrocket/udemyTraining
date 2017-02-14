package main

import (
	"fmt"
	//"time"
)

const p string = "death and taxes"

func main() {
	myInt := 10
	fmt.Println("Program Start....")
	makeCall()
	fmt.Printf("%d \t %v \n", &myInt, &myInt)
	fmt.Println(&myInt)
}

func makeCall() {

	const q = 42
	fmt.Println("q - ", q)
	fmt.Println("p - ", p)

	type MyBool bool
	const True = true
	const TypedTrue bool = true
	var mb MyBool
	mb = true
	mb = True
	fmt.Println(mb)

}
