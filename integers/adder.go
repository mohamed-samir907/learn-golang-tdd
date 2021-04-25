package integers

import "fmt"

// Add two integers and return the sum of them.
func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(2, 3)

	fmt.Printf("got %d\n", sum)
}
