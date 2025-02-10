package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)

	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	m["one"] = 1
	m["two"] = 2

	fmt.Println(m, m1)

	v, ok := m["three"]
	fmt.Println(v, ok)

	m["fore"] = 4
	delete(m, "one")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
