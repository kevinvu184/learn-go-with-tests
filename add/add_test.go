package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdd() {
	sum := add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
