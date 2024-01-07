package test

import (
	"fmt"
	"testing"
)

func Multiple(A, B int) int {
	return A * B
}
func TestMultiple(t *testing.T) {
	var X, Y, Result = 2, 2, 4

	testResult := Multiple(X, Y)
	if testResult != Result {
		t.Errorf("%d != %d", Result, testResult)
	}

	t.Run("first", func(t *testing.T) {
		var X, Y, Result = 200, 2, 400

		testResult := Multiple(X, Y)
		if testResult != Result {
			t.Errorf("%d != %d", Result, testResult)
		} else {
			fmt.Println("first")
		}
	})
	t.Run("second", func(t *testing.T) {
		var X, Y, Result = 400, 2, 800

		testResult := Multiple(X, Y)
		if testResult != Result {
			t.Errorf("%d != %d", Result, testResult)
		} else {
			fmt.Println("second")
		}
	})
}
