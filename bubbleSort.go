package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//read input from user
func readInput() []int {
	var stdin *bufio.Reader
	stdin = bufio.NewReader(os.Stdin)
	fmt.Println("Insert some numbers for your array: (ex. 1,2,3)")
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

func bubbleSort(array []int) []int {
	flag := true
	n := len(array) - 2
	for flag {
		flag = false
		for i := 0; i <= n; i++ {
			if array[i] > array[i+1] {
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
				flag = true
			}
		}
		n = n - 1
	}
	return array
}

func main() {
	intArray := readInput()
	orderedArray := bubbleSort(intArray)
	fmt.Printf("The ordered array is %v\n", orderedArray)
}
