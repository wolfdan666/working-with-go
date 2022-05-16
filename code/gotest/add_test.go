package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd1(t *testing.T) {
	sum := Add(1, 1)
	if sum != 2 {
		t.Errorf("Expected %d => Received %d", 2, sum)
	}
}

func TestAdd2(t *testing.T) {
	var testData = []struct {
		param1 int
		param2 int
		expect int
	}{
		{1, 2, 3},
		{0, 1, 1},
		{1, 0, 1},
		{-1, 1, 0},
		{-1, -1, -2},
	}

	for _, td := range testData {
		result := Add(td.param1, td.param2)
		if td.expect != result {
			t.Errorf("Expected %d but received %d", td.expect, result)
		}
	}
}
