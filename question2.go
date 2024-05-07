package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func increse(num int) int {
	i := 1
	for {
		if i <= num {
			num += i
		} else {
			break
		}
		i *= 10
	}
	return num
}

func decrease(num int) int {
	i := 1
	for {
		if i <= num {
			num -= i
		} else {
			break
		}
		i *= 10
	}
	return num
}

func decodeAndMinimize(encoded string) string {
	
	pairs := make([]string, 0)
	for i := 0; i < len(encoded); i++ {
		pairs = append(pairs, encoded[i:i+1])
	}
	
	result := 0
	for _, pair := range pairs {
		if strings.Contains(pair, "L") {
			if result == 0 {
				result = 21
			} else {
				digit := result % 10
				if digit == 0 {
					result = increse(result) * 10 + 0
				} else {
					result = result * 10 + 0
				}
			}
		} else if strings.Contains(pair, "R") {
			if result == 0 {
				result = 1
			} else {
				digit := result % 10
				result = result * 10 + digit + 1
			}

		} else {
			if result == 0 {
				result = 0
			} else {
				digit := result % 10
				result = result * 10 + digit
			}
		}
	}
	for{
		numStr := strconv.Itoa(result)
		if strings.Contains(numStr, "0") {
			break
		} else {
			result = decrease(result)
		}
	}
	resultString := strconv.Itoa(result)
	first := pairs[0]
	if len(resultString) != len(pairs) + 1 {
		count := len(pairs) - len(resultString) + 1
		for i := 0; i < count; i++ {
			resultString = string("0") + resultString	
		}
	}
	if strings.Contains(first, "=") {
		for i, pair := range pairs {
			if strings.Contains(pair, "=") {
				runes := []rune(resultString)
				runes[i] = rune(resultString[i+1])
				resultString = string(runes)
			} else {
				break
			}
		}
	}
	return resultString
}

func main() {
	fmt.Println("Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a encoded: ")
		scanner.Scan()
		encoded := scanner.Text()

		if strings.ToLower(encoded) == "exit" {
			fmt.Println("Exiting...")
			break
		}

		decoded := decodeAndMinimize(encoded)
		fmt.Println("Output a decoded :", decoded)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}