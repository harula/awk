package main

import "fmt"

func romanToInt(s string) int {

	var m = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(s); i++ {
		value := m[s[i]]
		if i < len(s)-1 && value < m[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}
	return total

}

func main() {
	testCase := []string{"III", "IV", "VI"}
	for _, s := range testCase {
		total := romanToInt(s)
		fmt.Println("total:", total)
	}
}
