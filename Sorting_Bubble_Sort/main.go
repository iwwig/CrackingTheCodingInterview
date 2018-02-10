package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func run(cnts string, data string) (int, int64, int64) {
	cnts = strings.Trim(cnts, "\n")
	cnt, _ := strconv.ParseInt(cnts, 10, 32)
	resultSlice := stringToIntArray(data, " ")
	swaps := bubbleSort(int(cnt), resultSlice)

	return swaps, resultSlice[0], resultSlice[len(resultSlice)-1]
}


func stringToIntArray(data string, sep string)  []int64    {
	// преобразуем data в массив
	data = strings.Trim(data, "\n")
	result := strings.Split(data, sep)

	var to []int64

	for r := range result {
		val, _ := strconv.ParseInt(result[r], 10, 32)
		to = append(to, val)
	}
	return to[:]
}

func bubbleSort(cntr int, result []int64) int {
	count := 0
	for r := range result {
		current := result[r]
		if cntr-1 > r {
			next := result[r+1]
			if current > next {
				count++
				result[r] = next
				result[r+1] = current
				cnt := bubbleSort(cntr, result)
				count = count + cnt
			}
		}
	}
	return count
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	// reader
	reader := bufio.NewReader(os.Stdin)

	cnt, _ := reader.ReadString('\n')
	data, _ := reader.ReadString('\n')

	swaps, min, max := run(cnt, data)

	fmt.Println(fmt.Sprintf("Array is sorted in %d swaps.", swaps))
	fmt.Println(fmt.Sprintf("First Element: %d", min))
	fmt.Println(fmt.Sprintf("Last Element: %d", max))
}