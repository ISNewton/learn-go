package main

import "fmt"

func main() {
	myArray := []int32{5, 6}

	myArray[0] = 2323

	mySlice := append(myArray, 7)

	fmt.Println(mySlice, cap(mySlice), len(mySlice))
}
