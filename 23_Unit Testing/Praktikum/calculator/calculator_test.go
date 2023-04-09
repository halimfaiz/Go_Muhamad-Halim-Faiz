package calculator

import (
	"testing"
)

func TestAdditon(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Addition(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}
func TestSubtraction(t *testing.T) {
	if Subtraction(1, 2) != -1 {
		t.Error("Expected 1 (-) 2 to equal -1")
	}
	if Subtraction(-1, -2) != 1 {
		t.Error("Expected -1 (-) -2 to equal 1")
	}
}
func TestDivision(t *testing.T) {
	if Division(10, 2) != 5 {
		t.Error("Expected 10 (/) 2 to equal to 5")
	}
	if Division(-10, -2) != 5 {
		t.Error("Expected -10 (/) -2 to equal to 5")
	}
	if Division(3, 2) != "1.50" {
		t.Error("Expected 3 (/) 2 to equal to 1.50")
	}
	if Division(-3, -2) != "1.50" {
		t.Error("Expected -3 (/) -2 to equal to 1.50")
	}
	if Division(0, 1) != 0 {
		t.Error("Expected 0 (/) 1 to equal to 0")
	}
}
func TestMultiplication(t *testing.T) {
	if Multiplication(1, 2) != 2 {
		t.Error("Expecter 1 (*) 2 to equal to 2")
	}
	if Multiplication(-1, -2) != 2 {
		t.Error("Expecter -1 (*) -2 to equal to 2")
	}
	if Multiplication(0, 1) != 0 {
		t.Error("Expecter 0 (*) 1 to equal to 0")
	}
	if Multiplication(1, 0) != 0 {
		t.Error("Expecter 1 (*) 0 to equal to 0")
	}
}
