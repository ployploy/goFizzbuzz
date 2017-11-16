package Test

import (
	"testing"
)

func TestInput1_output1(t *testing.T) {
	fizzbuzz := FizzBuzz{}
	result := fizzbuzz.show(1)
	expected := "1"
	if result != expected {
		t.Fatal("false")
	}
}
