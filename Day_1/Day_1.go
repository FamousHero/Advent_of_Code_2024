package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math"
	"strconv"
)
// Read the file, split each line, read the two numbers and 
// get abs distance & add it to running sum, return sum 
func main() {
	fi, err := os.Open("Day_1_input.txt")
	if err != nil{
		panic(err)
	}
	defer func(){
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	
	dataScanner := bufio.NewScanner(fi)
	
	var data []string

	var leftnums []float64
	var rightnums []float64

	for dataScanner.Scan() {
		data = append(data, dataScanner.Text())
	} 
	var sum float64 = 0
	for _, line := range data{
		fmt.Println("---Line---")
		nums := strings.Split(line, "   ");
		
		num1, _ := strconv.ParseFloat(nums[0], 64);
		append(leftnums, num1)

		num2, _ := strconv.ParseFloat(nums[1], 64);
		append(rightnums, num2)

		diff := math.Abs(num1 - num2);
		sum += diff
		// fmt.Println(nums)
		// fmt.Println(num2)
		fmt.Println("Diff: " + strconv.FormatFloat(diff, 'f', -1, 64))
		fmt.Println("Current Sum: " + strconv.FormatFloat(sum, 'f', -1, 64))
		fmt.Println("----------")
	}
	fmt.Println(len(data))

}
