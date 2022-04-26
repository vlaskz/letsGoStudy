package linear

import "fmt"

//time for this operation is linear, so is O(n)
func Example() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}
}
