package main

import (
	"fmt"
)

func main() {

	// Wrong Logic,
	x_wrong := []string{"a", "b", "c", "d"}
	y_wrong := x_wrong[:2]
	fmt.Println(cap(x_wrong), cap(y_wrong))
	y_wrong = append(y_wrong, "z")
	fmt.Println("x:", x_wrong)
	fmt.Println("y:", y_wrong)

	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
