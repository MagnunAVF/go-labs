package main

import "testing"

// The Bad: Repetitive, Brittle Tests If you have 10 cases, you'll have 10 nearly
//  identical functions. It's hard to see the "business rules" through the boilerplate.
// func TestAddPositive(t *testing.T) {
//     if Add(1, 2) != 3 {
//         t.Error("Expected 3")
//     }
// }

// func TestAddNegative(t *testing.T) {
//     if Add(-1, -1) != -2 {
//         t.Error("Expected -2")
//     }
// }

// The Good: Table-Driven Tests with Subtests Separates the test data (The "What")
// from the test execution (The "How").
func Add(a, b int) int {
	return a + b
}
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -1, -2},
		{"mixed numbers", -1, 5, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
