package Test

import "strconv"

type FizzBuzz struct {
}

func (FizzBuzz) show(number int) string {
	isFizzBuzz := number%3 == 0 && number%5 == 0
	if isFizzBuzz {
		return "FizzBuzz"
	}
	if isFizz(number) {
		return "Fizz"
	}
	if isBuzz(number) {
		return "Buzz"
	}
	return strconv.Itoa(number)
}

func isFizz(number int) bool {
	return number%3 == 0
}

func isBuzz(number int) bool {
	return number%5 == 0
}
