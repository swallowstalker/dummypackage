package dummypackage

import "fmt"

func Echo(s string) string {
	fmt.Println("Echo print this string: " + s)
	return "Echo print this string: " + s
}

func Print(s string) {
	fmt.Println("Print result: " + s)
}

func Print2(s string) {
	fmt.Println("Print2 result: " + s)
}

func Print3(s string) {
	fmt.Println("Print3 result: " + s)
}
