package main

import "fmt"

func Frequency(s string) map[string]int {
	result := make(map[string]int)

	for _, ch := range s {
		result[string(ch)]++
	}

	return result
}

func main() {
	fmt.Println(Frequency("MalayAlam"))
}
