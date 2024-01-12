// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() {
// 	fmt.Sprintf("My name is %s %d tuổi", p.Name, p.Age)
// 	// fmt.Println(a)
// 	// return fmt.Sprintf("My name is %s %d tuổi", p.Name, p.Age)
// }

// func main() {
// 	// myProfile := Person{"Linh", 24}
// 	// a := fmt.Sprintf("My name is %s %d tuổi", myProfile.Name, myProfile.Age)
// 	// fmt.Println(a)
// }

package main

import "fmt"

func sayHello(name string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}

func main() {
	// Goroutine
	go sayHello("Viet")

	// normal function
	sayHello("Nam")
}
