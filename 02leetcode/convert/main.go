package main

import "fmt"

func convert(s string, numRows int) string {

	if numRows == 1 || len(s) <= numRows {
		return s
	}
	var strs = make([]string, min(numRows, len(s)))
	row := 0
	goingDown := false
	for _, c := range s {
		strs[row] += string(c)
		if row == 0 || row == numRows-1 {
			//每首列或最后一列变换上下行方向
			goingDown = !goingDown
		}
		if goingDown {
			row++
		} else {
			row--
		}

	}

	rst := ""
	for _, s := range strs {
		rst += s
	}
	return rst
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	input := "PAYPALISHIRING"
	numRows := 3
	output := convert(input, numRows)
	fmt.Println(output) // Output: PAHNAPLSIIGYIR
}
