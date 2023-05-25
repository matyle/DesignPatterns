package ch1

import "testing"

func TestRun(t *testing.T) {
	r := Run(1, 2, "+")
	if r != 3 {
		t.Error("1 + 2 != 3")
	}
	r = Run(1, 2, "-")
	if r != -1 {
		t.Error("1 - 2 != -1")
	}
}
