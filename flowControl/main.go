package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 121
	m := strconv.Itoa(i)
	for _, v := range m {
		fmt.Println(v)
	}
	x := make([]byte, len(m))
	for i := len(x) - 1; i > 0; i-- {
		x = append(x, m[i])
	}
	fmt.Println(x)
}

//попытка решить задачу с литкода)
