package main

import (
	"testing"
	"fmt"
	"os"
	"bufio"
)

func TestRun(t *testing.T) {
	result := run("abcccc", "cde")
	if (result != 7) {
		fmt.Print(result)
		t.Fail()
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestRunFromFile(t *testing.T) {
	// reader
	//dat, err := ioutil.ReadFile("input/input10.txt")
	f, err := os.Open("input/input10.txt")
	check(err)

	reader := bufio.NewReader(f)

	fmt.Println("Enter first: ")
	first, _ := reader.ReadString('\n')
	//fmt.Println(first)

	fmt.Println("Enter second: ")
	second, _ := reader.ReadString('\n')

	result := run(first, second)
	if result != 734 {
		fmt.Print(result)
		t.Fail()
	}
}
