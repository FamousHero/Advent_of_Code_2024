package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"math"

)

func Day2(scanner *bufio.Scanner) {
	safe_count := 0
	for scanner.Scan(){
		numbers := strings.Split(scanner.Text(), " ")
		number_len := len(numbers)
		decreasing_order := false
		prev, _ := strconv.ParseInt(numbers[0], 10, 64)
		for i:= 1; i < number_len; i++{
			x, _ := strconv.ParseInt(numbers[i], 10, 64)
			
			diff := x - prev
			prev = x
			
			if i == 1 && diff < 0{ // decreasing order
				decreasing_order = true
			}
			// Check conditions for safe
			if diff == 0 || (diff < 0 && !decreasing_order) ||
			(diff > 0 && decreasing_order) || math.Abs(float64(diff)) > 3{
				break
			}
			if i == number_len-1{
				safe_count += 1
			}
		}
	}
	fmt.Println(safe_count)
}

func main(){
	file, err := os.Open("Day_2_input.txt")
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

