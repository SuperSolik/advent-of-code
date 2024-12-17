package main

import "testing"

func TestCase1(t *testing.T) {
	c := Computer{A: 0, B: 0, C: 9, program: []int{2, 6}}

	c.run()

	if c.B != 1 {
		t.Fatalf("c: %v\nB != 1\n", c)
	}
}

func TestCase2(t *testing.T) {
	c := Computer{A: 10, B: 0, C: 0, program: []int{5, 0, 5, 1, 5, 4}}

	c.run()

	expected := []int{0, 1, 2}

	if len(c.out) != len(expected) {
		t.Fatalf("out = %v is not expected %v\n", c.out, expected)
	}

	for i := 0; i < len(c.out); i++ {
		if c.out[i] != expected[i] {
			t.Fatalf("out = %v is not expected %v\n", c.out, expected)
		}
	}
}

func TestCase3(t *testing.T) {
	c := Computer{A: 2024, B: 0, C: 0, program: []int{0, 1, 5, 4, 3, 0}}

	c.run()

	expected := []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}

	if len(c.out) != len(expected) {
		t.Fatalf("out = %v is not expected %v\n", c.out, expected)
	}

	for i := 0; i < len(c.out); i++ {
		if c.out[i] != expected[i] {
			t.Fatalf("out = %v is not expected %v\n", c.out, expected)
		}
	}

	if c.A != 0 {
		t.Fatalf("c: %v\nB != 1\n", c)
	}
}

func TestCase4(t *testing.T) {
	c := Computer{A: 0, B: 29, C: 0, program: []int{1, 7}}

	c.run()

	if c.B != 26 {
		t.Fatalf("c: %v\nB != 1\n", c)
	}
}

func TestCase5(t *testing.T) {
	c := Computer{A: 0, B: 2024, C: 43690, program: []int{4, 0}}

	c.run()

	if c.B != 44354 {
		t.Fatalf("c: %v\nB != 1\n", c)
	}
}
