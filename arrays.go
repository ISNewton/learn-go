
package arrays

func getMyNumber(theNumber int, secondN int) (int, int, error) {
	var err error
	if theNumber > secondN {
		err = errors.New("Invalid number")
		return theNumber, secondN, err
	}
	return theNumber + 1, secondN, err
}

