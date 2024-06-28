package main

func Addition(a, b int) int {
	return *Subtraction(a, b) + 2*a
}

func Subtraction(a, b int) *int {
	c := a - b
	return &c
}
