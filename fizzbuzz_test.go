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

func TestInput4_output4(t *testing.T) {
	fizzbuzz := FizzBuzz{}
	result := fizzbuzz.show(4)
	expected := "4"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInput5_outputBuzz(t *testing.T) {
	fizzbuzz := FizzBuzz{}
	result := fizzbuzz.show(5)
	expected := "Buzz"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInput6_outputFizz(t *testing.T) {
	fizzbuzz := FizzBuzz{}
	result := fizzbuzz.show(6)
	expected := "Fizz"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}
