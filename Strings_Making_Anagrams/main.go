package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
)


func run(first string, second string) (int) {
	//составим массивы букв по каждому вхождению


	second = strings.Replace(second, "\n", "", -1)
	first = strings.Replace(first, "\n", "", -1)
	// составим массив символов и частотность
	// получим массив символов от a до z
	alfabet := map[string]int{"a":0,"b":0,"c":0,"d":0,"e":0,"f":0,"g":0,"h":0,"i":0,"j":0,"k":0,"l":0,"m":0,"n":0,"o":0,"p":0,"q":0,"r":0,"s":0,"t":0,"u":0,"v":0,"w":0,"x":0,"y":0,"z":0}
	result_f := make(map[string]int32)
	result_s := make(map[string]int32)
	for r := range first {
		result_f[string(first[r])]++
	}

	for r := range second {
		result_s[string(second[r])]++
	}

	// переберем по алфавиту
	cnt := 0
	for r := range alfabet {
		// посчитаем разницу по модулю
		cnt = cnt + int(math.Abs(float64(result_f[r]) - float64(result_s[r])))

	}

	return cnt
}


func main() {

	// reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first: ")
	first, _ := reader.ReadString('\n')
	//fmt.Println(first)

	fmt.Println("Enter second: ")
	second, _ := reader.ReadString('\n')

	run(first, second)
}
