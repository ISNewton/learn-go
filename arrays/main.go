package main

import "fmt"

func main() {
	myArray := []int32{5, 6}

	myArray[0] = 2323

	mySlice := append(myArray, 7)

	fmt.Println(mySlice, cap(mySlice), len(mySlice))

	myMap := map[int8]string{8: "Ash"}

	ash, ok := myMap[8]

	if ok {
		fmt.Println("We found ash at :", ash)
		delete(myMap, 8)

		fmt.Println("We deleted ash")
	} else {
		fmt.Println("Something went wrong , we can't find a map")
	}
}
