package dummypackage

import "fmt"

func Echo(s string) string {
	fmt.Println("Echo print this string: " + s)
	return "Echo print this string: " + s
}
