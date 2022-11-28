package calculator_test

import (
	"testing"

	"github.com/ReyAdrian520/PivotTech/calculator"
)

func TestAdd(t *testing.T) {
	got := calculator.Add(2, 1)
	expected := 3
	if got != expected {
		t.Errorf("not expected results, got: %v expected: %v", got, expected)
	}
}

func TestSub(t *testing.T) {
	got := calculator.Sub(2, 1)
	expected := 1
	if got != expected {
		t.Errorf("not expected results, got: %v expected: %v", got, expected)
	}
}

func TestMuti(t *testing.T) {
	got := calculator.Muti(2, 1)
	expected := 2
	if got != expected {
		t.Errorf("not expected results, got: %v expected: %v", got, expected)
	}
}
func TestDiv(t *testing.T) {
	result, err := calculator.Div(6, 3)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if result != 2 {
		t.Error("Expected 2, got ", result)
	}
}

func TestSubTable(t *testing.T) {
	data := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 2, y: 1, expected: 1},
		{x: 2, y: 2, expected: 0},
		{x: 2, y: 3, expected: -1},
	}
	for _, val := range data {
		got := calculator.Sub(val.x, val.y)
		if got != val.expected {
			t.Errorf("for x: %v and y: %v expected: %v but got: %v", val.x, val.y, val.expected, got)
		}
	}

}
