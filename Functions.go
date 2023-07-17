// package main

// import "fmt"

// func main() {
// 	add(4, 5)
// 	add(2, 9)
// 	add(43, 5)
// }

// func add(x int, y int) {
// 	var z = x + y
// 	fmt.Println("x + y =", z)
// }

// package main

// import "fmt"

// func main() {
// 	add(1, 2, 3.4, 5.6, 1.2)
// }
// func add(x, y int, a, b, c float32) {
// 	var z = x + y
// 	var d = a + b + c
// 	fmt.Println("x + y = ", z)
// 	fmt.Println("a + b + c = ", d)
// }

// package main

// import "fmt"

// func main() {
// 	var a = 8
// 	fmt.Println("a before", a)
// 	increment(a)
// 	fmt.Println("a after", a)
// }

// func increment(x int) {
// 	fmt.Println("x before", x)
// 	x = x + 20
// 	fmt.Println("x afner", x)

// }

package main

import "fmt"

func main() {
	add(1, 2, 3)       // sum = 6
	add(1, 2, 3, 4)    // sum = 10
	add(5, 6, 7, 2, 3) // sum = 23
}

func add(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum=", sum)
}
