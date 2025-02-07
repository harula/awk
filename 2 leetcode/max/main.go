package main

var data = []struct {
	id uint
	typeN enum('0','1')
	opAt int
}{
	{1,'0',123456},{2,'0',23456,},{3,'1',23457},{4,'1',23458}
}

func maxCount(data []struct{id uint
	typeN enum('0','1')
	opAt int})int{
		var rst int
		for _,d := range data{
			if d.typeN == '0'{
				maxCount ++
			}else if d.typeN == '1' {
				maxCount --
			}
			if maxCount > rst {
				rst = maxCount
			}
		}
		return rst
	}