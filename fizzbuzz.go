package Test

import "strconv"

type FizzBuzz struct {
}

func (FizzBuzz) show(number int) string {
	if number == 3 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
