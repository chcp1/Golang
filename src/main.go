package main

import "fmt"

func main() {
	var A string
	fmt.Scan(&A)
	for i := 0; i < len(A); i++ {
		fmt.Print(string(A[i]))
	}
	fmt.Print("\n")

}
