package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read an array of numbers
	length := 0
	fmt.Println("Enter the number of inputs")
	fmt.Scanln(&length)
	fmt.Println("Enter the inputs")
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}
	fmt.Println(numbers)

	// Read a string
	fmt.Println("Enter a string")
	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')
	fmt.Println(line)
}
