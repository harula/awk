package main

import "fmt"

func intToRoman2(val int) string {
	var rst string
	vs := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ss := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i, v := range vs {
		for val >= v {
			rst += ss[i]
			val -= v
		}

	}
	return rst
}

func intToRoman(val int) string {
	var rst string
	for val > 0 {
		if val >= 1000 {
			rst += "M"
			val -= 1000
		} else if val >= 900 {
			rst += "CM"
			val -= 900
		} else if val >= 500 {
			rst += "D"
			val -= 500
		} else if val >= 400 {
			rst += "CD"
			val -= 400
		} else if val >= 100 {
			rst += "C"
			val -= 100
		} else if val >= 90 {
			rst += "XC"
			val -= 90
		} else if val >= 50 {
			rst += "L"
			val -= 50
		} else if val >= 40 {
			rst += "XL"
			val -= 40
		} else if val >= 10 {
			rst += "X"
			val -= 10
		} else if val >= 9 {
			rst += "IX"
			val -= 9
		} else if val >= 5 {
			rst += "V"
			val -= 5
		} else if val >= 4 {
			rst += "IV"
			val -= 4
		} else if val >= 1 {
			rst += "I"
			val -= 1
		}

	}
	return rst
}

func main() {
	testCases := []int{3749, 58}
	for _, n := range testCases {
		s := intToRoman2(n)
		fmt.Println("s:", s)
	}
}
