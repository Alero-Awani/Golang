package main

// import "fmt"

// import "fmt"

// //EXAMPLE 1 - NOTE that if you dont specify the reciver as a pointer, it still works as long as you are using a return statement with the method
// // in the struct
// type Rectangle struct {
// 	length  int
// 	breadth int
// }

// func (r Rectangle) area() int {
// 	return r.length * r.breadth
// }

// func main() {
// 	r := Rectangle{breadth: 10, length: 5}
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// }

// //EXAMPLE 2
// type Rectangle struct {
// 	length  int
// 	breadth int
// }

// func (r Rectangle) area() int {
// 	return r.length * r.breadth
// }

// func (r *Rectangle) incLength(n int) {
// 	for i := 0; i < n; i++ {
// 		r.length += i
// 	}
// }

// func main() {
// 	r := Rectangle{breadth: 10, length: 5}
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// 	r.incLength(7)
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// }

//EXAMPLE 3 - HERE WE CAN SEE THAT WE CAN CREATE A SLICE OF MAPS USING A STRUCT
// type Employee struct {
// 	eid int
// 	id  int
// }

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 			employees[i] = Employee{i, i + 10}
// 			fmt.Println(employees[i])
// 	}
// }

// //EXAMPLE 4 -
// type Employee struct {
// 	eid int
// 	id  int
// }

// func (e Employee) get_id() int {
// 	return e.eid + 10
// }

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 			employees[i] = Employee{eid: i}
// 			employees[i].id = employees[i].get_id()
// 			fmt.Printf("%+v\n", employees[i])
// 	}
// }