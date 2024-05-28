package main

import "fmt"

type car struct {
	name        string
	model       uint16
	isAvailable bool
}

type carWithOwner struct {
	ownerName string
	hisCar    car
}

func (c carWithOwner) wheels() uint8 {
	return 1
}

type somethingCanGo interface {
	wheels() uint8
}

func main() {

	someBody := carWithOwner{ownerName: "tr"}

	fmt.Println(canI(someBody))

}

func canI(someCar somethingCanGo) string {
	var theWheels = someCar.wheels()
	if theWheels >= 4 {
		return "Yes you can."
	} else if theWheels < 4 && theWheels > 1 {
		return "I doubt it."
	} else {
		return "Oops! you can not."
	}
}
