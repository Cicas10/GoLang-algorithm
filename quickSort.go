package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// quick sort algorith
func quickSort(list []int, start int, end int) []int {
	if start < end {
		// partition the list
		split := partition(list, start, end)
		// sort both halves
		quickSort(list, start, split-1)
		quickSort(list, split+1, end)
	}
	return list
}

func partition(list []int, start int, end int) int {
	pivot := list[start]
	left := start + 1
	right := end
	done := false
	for !done {
		for left <= right && list[left] <= pivot {
			left = left + 1
		}
		for list[right] >= pivot && right >= left {
			right = right - 1
		}

		if right < left {
			done = true
		} else {
			temp := list[left]
			list[left] = list[right]
			list[right] = temp
		}
	}
	// swap start with list[right]
	temp := list[start]
	list[start] = list[right]
	list[right] = temp
	return right
}

//read input from user and parse string to int
func readInput() []int {
	var stdin *bufio.Reader
	stdin = bufio.NewReader(os.Stdin)
	fmt.Println("Insert some numbers to be ordered: (ex. 3,1,2)")
	s, _ := stdin.ReadString('\n')
	s = s[0 : len(s)-1]
	//testString := "2,5,13,59"
	stringArray := strings.Split(s, ",")
	//fmt.Printf("stringArray is %v\n", stringArray)
	intArray := []int{}
	for _, i := range stringArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intArray = append(intArray, j)
	}
	return intArray
}

func main() {
	intArray := readInput()
	orderedArray := quickSort(intArray, 0, len(intArray)-1)
	fmt.Printf("The ordered array is %v\n", orderedArray)
}
