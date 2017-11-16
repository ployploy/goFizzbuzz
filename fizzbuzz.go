package Test

type FizzBuzz struct {
}

func (FizzBuzz) show(number int) string {
	if number == 2 {
		return "2"
	}
	return "1"
}
