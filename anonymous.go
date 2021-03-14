package Anonymous

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("Anonymous function")
	}()

	func(x int) {
		fmt.Println("Anonymous function passing arguments", x)
	}(42)
}

func foo() {
	fmt.Println("In foo func")
}
