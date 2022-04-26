package quadratic

import "fmt"

//time for this operation is quadratic inner 'for' iterates n times while outer 'for' iterates once, so is O(n^2)
func Example() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("counting from 0 to i*10+j  = ", i*10+j)
		}
	}
}
