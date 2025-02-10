package main

type LogData struct {
	id      int
	isLogin bool
	opAt    int
}

var data = []LogData{
	{1, true, 123456}, {2, true, 23456}, {3, false, 23457}, {4, false, 23458},
}

var maxLogin int

func max(data []LogData) int {
	var rst int
	for _, d := range data {
		if d.isLogin {
			maxLogin++
		} else {
			maxLogin--
		}
		if maxLogin > rst {
			rst = maxLogin
		}
	}
	return rst
}
