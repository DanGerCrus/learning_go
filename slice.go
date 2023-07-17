// package main

// import "fmt"

// var users = []string{"Tom", "Alice", "Kate"}

// func main() {
// 	fmt.Println(users[2]) // Kate
// 	users[2] = "Katherine"

// 	for _, value := range users {
// 		fmt.Println(value)
// 	}

// 	users := []string{"Tom", "Alice", "Kate"}
// 	users = append(users, "dggergkihdjgbjgjudfgb5hrkirtegvhj")

// 	for _, value := range users {
// 		fmt.Println(value)
// 	}

// }

// package main

// import "fmt"

// func main() {
// 	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
// 	users1 := initialUsers[2:6] // с 3-го по 6-й
// 	users2 := initialUsers[:4]  // с 1-го по 4-й
// 	users3 := initialUsers[3:]  // с 4-го до конца

// 	fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
// 	fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
// 	fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]
// }

package main

import "fmt"

func main() {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	//удаляем 4-й элемент
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]
}
