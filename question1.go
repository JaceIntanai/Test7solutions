package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func maxPathNodeSum(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i])
	}

	// Reverse from bottom to top
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] += max(dp[i+1][j], dp[i+1][j+1])
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	// Read input from JSON file
	file, err := ioutil.ReadFile("hard.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Parse JSON data
	var data [][]int
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	result := maxPathNodeSum(data)
	fmt.Println("Maximum path sum:", result)
}
