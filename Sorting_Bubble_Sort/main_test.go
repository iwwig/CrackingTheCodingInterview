package main

import (
	"testing"
	"os"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestRunFromFile(t *testing.T) {
	// reader
	f, err := os.Open("input/input1.txt")
	check(err)

	reader := bufio.NewReader(f)

	cnt, _ := reader.ReadString('\n')
	data, _ := reader.ReadString('\n')

	swaps, min, max := run(cnt, data)
	if swaps != 68472  || min != 4842 || max != 1994569 {
		t.Fail()
	}
}
