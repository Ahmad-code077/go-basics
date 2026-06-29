package main

import "fmt"

func main() {

	numbers := [4]int{12, 20, 24, 15}

	highestMark := numbers[0]
	highestMarkIndex := 0

	for i, data := range numbers {
		if data > highestMark {
			highestMark = data
			highestMarkIndex = i
		}
	}

	fmt.Println("First Student to check", highestMark, highestMarkIndex)
}
