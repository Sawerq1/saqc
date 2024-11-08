package main

import "fmt"

func Printfln(temlate string, values ...interface{}) {
	fmt.Printf(temlate+"\n", values...)
}
