package main

import "testing"

func TestZero(t *testing.T) {
	result := fibonnaci(0)
	if result != 0 {
		t.Errorf("fibonnaci(0) = %d; want 0", result)
	}
}

func TestOne(t *testing.T) {
	result := fibonnaci(1)
	if result != 1 {
		t.Errorf("fibonnaci(1) = %d; want 1", result)
	}
}

func TestTwo(t *testing.T) {
	result := fibonnaci(2)
	if result != 1 {
		t.Errorf("fibonnaci(1) = %d; want 1", result)
	}
}

func TestThree(t *testing.T) {
	result := fibonnaci(3)
	if result != 2 {
		t.Errorf("fibonnaci(2) = %d; want 2", result)
	}
}

func TestFour(t *testing.T) {
	result := fibonnaci(4)
	if result != 3 {
		t.Errorf("fibonnaci(4) = %d; want 3", result)
	}
}
