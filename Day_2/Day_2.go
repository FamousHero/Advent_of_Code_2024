package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	// "strconv"

)

func Day2(scanner *bufio.Scanner) {
	for scanner.Scan(){
		words := strings.Split(scanner.Text(), " ")
		for i, word := range words{
			var pyramid string;
			// String concatenation should be done with strings.Builder or strings.Join
			for ;i > -1; i--{
				pyramid = pyramid + "-"
			}
			pyramid = pyramid + " "
			fmt.Println(pyramid + word)
		}
		fmt.Println(words)
	}
}

func main(){
	file, err := os.Open("test.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	Day2(scanner)

	if err := scanner.Err(); err != nil{
		panic(err)
	}
}

