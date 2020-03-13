package main

import "testing"

func TestTen(t *testing.T) {
	result := threeAndFive(10)
	if result != 23 {
		t.Errorf("threeAndFive(10) = %d; want 23", result)
	}
}

func TestOneThousand(t *testing.T) {
	result := threeAndFive(1000)
	if result != 233168 {
		t.Errorf("threeAndFive(1000) = %d; want 233168", result)
	}
}
