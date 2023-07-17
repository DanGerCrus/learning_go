// package main

// import "fmt"

// func main() {
// 	var a = add(4, 5)  // 9
// 	var b = add(20, 6) // 26
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

//	func add(x, y int) (z int) {
//		z = x + y
//		return
//	}
package main

import "fmt"

func main() {
	var age, name, xrust = add(4, 5, "Tom", "Simpson", "love", "Trop")
	fmt.Println(age)   // 9
	fmt.Println(name)  // Tom Simpson
	fmt.Println(xrust) // love Trop
}

func add(x, y int, firstName, lastName string, xru, stu string) (int, string, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	var love = xru + " " + stu
	return z, fullName, love
}
