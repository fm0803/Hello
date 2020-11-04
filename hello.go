package main

import (
	"fmt"
	//"github.com/fm0803/stringutil"
)

func main() {
	fmt.Println("!oG ,olleH")
	//fmt.Println(stringutil.Reverse("!oG ,olleH"))

	str2 := "hello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2[:])
	fmt.Println(str2)
	data2 = nil
	fmt.Println(data2)
}
