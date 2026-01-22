package object

import (
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "hi my name is john"}
	diff2 := &String{Value: "hi my name is john"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same value have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same value have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("string with different value have same hash key")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}
	two := &Integer{Value: 2}

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integers with same value have different hash keys")
	}

	if one1.HashKey() == two.HashKey() {
		t.Errorf("integers with different values have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("booleans with same value have different hash key")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("boolean with different values have same hash key")
	}
}
