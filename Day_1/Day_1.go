package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math"
	"strconv"
	"sort"
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
	var sum float64
	for _, line := range data{
		fmt.Println("---Line---")
		nums := strings.Split(line, "   ");
		
		num1, _ := strconv.ParseFloat(nums[0], 64);
		leftnums = append(leftnums, num1)

		num2, _ := strconv.ParseFloat(nums[1], 64);
		rightnums = append(rightnums, num2)

		diff := math.Abs(num1 - num2);
		sum += diff
		// fmt.Println(nums)
		// fmt.Println(num2)
		fmt.Println("Diff: " + strconv.FormatFloat(diff, 'f', -1, 64))
		fmt.Println("Current Sum: " + strconv.FormatFloat(sum, 'f', -1, 64))
		fmt.Println("----------")
	}
	sort.Float64s(leftnums)
	sort.Float64s(rightnums)
	
	sum = 0
	for count := 0; count < 1000; count++{
		sum += math.Abs(leftnums[count] - rightnums[count])
	}
	// fmt.Println(leftnums)
	// fmt.Println(sum)


	// ---------------- Part 2 -------------
	var occurrences = map[float64] int{} 
	for _, key := range rightnums{
		occurrences[key] += 1
	} 
	var similarity_score int
	for _, key := range leftnums{
		similarity_score += int(key) * occurrences[key]
	}
	// fmt.Println(similarity_score)

}
