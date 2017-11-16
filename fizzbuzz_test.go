package Test

import (
	"testing"
)

func TestInput1_output1(t *testing.T) {
	fizzbuzz := FizzBuzz{}
	result := fizzbuzz.show(1)
	expected := "1"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInput2_output2(t *testing.T) {
	fizzbuzz := FizzBuzz{}
	result := fizzbuzz.show(2)
	expected := "2"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInput3_outputFizz(t *testing.T) {
	fizzbuzz := FizzBuzz{}
	result := fizzbuzz.show(3)
	expected := "Fizz"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}
