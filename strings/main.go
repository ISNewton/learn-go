package main

import (
	"errors"
	"fmt"
)

func getMyNumber(theNumber int, secondN int) (int, int, error) {
	var err error
	if theNumber > secondN {
		err = errors.New("Invalid number")
		return theNumber, secondN, err
	}
	return theNumber + 1, secondN, err
}

func main() {

	var n1, n2, err = getMyNumber(4, 1)

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {

		fmt.Println(n1, n2)
	}

}
