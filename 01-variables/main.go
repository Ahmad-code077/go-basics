package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	shortVaribleDec := "This is short varible declecration"
	letter := 'B'

	fmt.Printf("%d %d %t %t %s %s %c\n",
		i,
		j,
		c,
		python,
		java,
		shortVaribleDec,
		letter,
	)
}
