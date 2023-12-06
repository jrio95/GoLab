package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	winnerElfCalories := 0
	currentElfCalories := 0

	for sc.Scan() {
		calLine, err := strconv.Atoi(sc.Text())
		currentElfCalories += calLine

		if err != nil {
			if currentElfCalories > winnerElfCalories {
				winnerElfCalories = currentElfCalories
			}
			currentElfCalories = 0
		}
	}

	finalTime := time.Since(startTime)
	fmt.Println(winnerElfCalories)
	fmt.Println("total time: ", finalTime)
}
