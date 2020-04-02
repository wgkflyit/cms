package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("hello cms")
	u1 := uuid.NewV4()
	fmt.Println(u1)
	fmt.Println(2 << 11)

	var f1 int64 = 12

	var f2 int64 = 5
	var f3 int64 = 5
	var f4 int64 = 5
	fmt.Println((f1 << 22) | f2<<12 | f3 | f4)

}
