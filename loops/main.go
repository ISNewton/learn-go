package main

import "fmt"

func main() {

	mySlice := []string{"Ash", "Abdo", "Alhaj"}

	for number := range mySlice {
		//fmt.Println(number, mySlice[number])
		if mySlice[number] == "alhaj" {
			//fmt.Println("We found Alhaj")
		}
	}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])

	}

	myMap := map[int]string{44: "Four Four"}

	for index, name := range myMap {
		fmt.Println(index, name)
	}

}
