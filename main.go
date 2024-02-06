package main

import "fmt"

// import "fmt"

// func main() {
// 	var name string = "Joe"
// 	var i int = 78
// 	fmt.Printf("Hey, %v! You have scored %d/100 in Physics", name, i)
//     // fmt.Println("Hello World")
// }

//CONSTANTS
// const PI float64 = 3.14

// func main() {
// 	var radius float64 = 5.154
// 	var area float64

// 	area = PI * radius * radius
// 	fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
// 	fmt.Printf("Area of circle is : %f", area)
// 	fmt.Println("Thank you")
// }

// When using scanf, if you dont specify the print statement before it, it will still allow you put in an input
// func main() {
// 	var name string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanf("%s", &name)
// 	fmt.Println("Hey there,", name)
// }

// MULTIPLE INPUTS WITH SCANF

// func main(){
// 	var name string
// 	var alpha_count int
// 	var float_value float32
// 	var bool_value bool

// 	fmt.Scanf("%s %d %g %t", &name, &alpha_count, &float_value, &bool_value)
// 	fmt.Printf("%s %d %g %t", name, alpha_count, float_value, bool_value)
// }

// func main() {
// 	var name string
// 	var is_muggle bool

// 	fmt.Print("Enter your name and are you a muggle?")
// 	count, err := fmt.Scanf("%s %t", &name, &is_muggle)
// 	fmt.Print(name, is_muggle)
// 	fmt.Println("Count: ", count)
// 	fmt.Println("error: ", err)
// }

// USING %T to get the type of variable

// func main() {
// 	var grades int = 42
// 	var message string = "hello world"
// 	var isCheck bool = true
// 	var amount float64 = 89.900

// 	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
// 	fmt.Printf("Variable message = '%v' is of type %T \n", message, message)
// 	fmt.Printf("Variable ischeck = '%v' is of type %T \n", isCheck, isCheck)
// 	fmt.Printf("Variable amount = %v is of type %T \n", amount, amount)
//

// USING reflect.Typeof to get the type of literal and variable

// LITERAL

// func main(){
// 	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
// }

//VARIABLE

// func main(){
// 	var grades int = 42
// 	fmt.Printf("Variable grades = %v is of type %v \n", grades, reflect.TypeOf(grades))
// }

// func main() {
//         var x, y int = 100,90
//         fmt.Println(x & y)
//         fmt.Println(x | y)
// }

// func main() {
//         var x, y int = 100,90
//         fmt.Println((x+y) >> 2)
// }

// func main() {
//         var x, y int = 100,90
//         fmt.Println(!(((x+y) >> 2 ) == 47))
// }

// func main(){
// 	var a string = "happy"
// 	if a == "happy" {
// 		fmt.Println(a)
// 	} else {
// 		fmt.Println("They are not the same thing ")
// 	}
// }

// func main(){
// 	var fruit string = "grapes"

// 	if fruit == "apples" {
// 		fmt.Println("The fruit is an apple")
// 	} else {
// 		fmt.Println("The fruit is not an apple")
// 	}
// }

// func main(){
// 	fruit:= "grapes"

// 	if fruit == "apples" {
// 		fmt.Println("The fruit is an apple")
// 	} else if fruit == "grapes" {
// 		fmt.Println("This fruit is a grape")
// 	} else {
// 		fmt.Println("What kind of fruit is this")
// 	}
// }

// func main() {
//         var a, b string = "kolkata", "Kolkata"
//         if a == b {
//                 fmt.Println("strings are equal")
//         } else {
//                 fmt.Println("strings are not equal")
//         }
//         fmt.Println("thank you!")

// }

//ARRAYS

// func main() {
// 	var fruits [2]string = [2]string{"apples", "oranges"}
// 	fmt.Println(fruits)

// 	marks := [3]int{10, 20, 30}
// 	fmt.Println(marks)

// 	names := [...]string{"Rachael", "phoebe", "Monica"}
// 	fmt.Println(names)
// }

//CHANGING THE VALUE OF AN ARRAY
// func main(){
// 	var grades [5]int = [5]int{10,20,30,40,50}
// 	grades[1] = 100
// 	fmt.Println(grades)
// }

//LOOPING THROUGH AN ARRAY
// func main(){
// 	var grades [5]int = [5]int{50, 60, 70, 80, 90}

// 	for i :=0; i < len(grades); i++ {
// 		fmt.Println(grades[i])
// 	}
// }

// SLICES
// func main(){
// 	grades := [3]int{10, 20, 30}
// 	slice := grades[1:]
// 	fmt.Println(slice)
// 	sub_slice := slice[:]
// 	fmt.Println(sub_slice)
// }

// Create slices from make function(note that the make function is usually used to create empty slices)

// func main() {
// 	slice := make([]int,5,8)
// 	fmt.Println(slice)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))
// }

// Appending to a slice(note that when you append to a slice, the orignal array also gets affected)

// func main() {
// 	arr := [5]int{10, 20, 30, 40, 50}
// 	slice := arr[1:3]

// 	fmt.Println(slice)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))

// 	slice = append(slice, 900, -90, 50)
// 	fmt.Println(slice)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))
// }

// CHANGING VALUE IN SLICE(when you change the value, the original array also gets affected)
// func main(){
// 	arr := [5]int{10, 20, 90, 70, 60}
// 	slice := arr[0:3]

// 	slice[2] = 900

// 	fmt.Println(arr)
// 	fmt.Println(slice)
// }

// APPENDING A SLICE TO ANOTHER SLICE
// WE CAN DO THIS USING THE THREE DOT OPERATOR

// func main(){
// 	arr := [5]int{10, 20, 30, 40, 50}
// 	slice := arr[:2]//10,20

// 	arr_2 := [5]int{5,15,25,35,45}
// 	slice_2 := arr_2[:2]

// 	new_slice := append(slice, slice_2...)

// 	fmt.Println(new_slice)
// }

//DELETING FROM A SLICE
//WE DO THIS BY CREATING A NEW SLICE THAT DOES NOT CONSIST OF THE ELEMENT THAT NEEDS TO BE DELETED

// func main() {
// 	arr := [5]int{10,20,30,40,50}// deleting 30
// 	i := 2
// 	fmt.Println(arr)

// 	slice_1 := arr[:i]//10,20
// 	slice_2 := arr[i+1:]//40,50

// 	new_slice := append(slice_1, slice_2...)

// 	fmt.Println(new_slice)
// }

func main(){
	arr := [5]int{10,20,30,40,50}
	i := 2//this deletes 20 and not 30, because in the todo app, 2 is actually 2 and the/when a person want to delete 2 in the 0 index they actually want to delte 2 -1 
	fmt.Println(arr)

	slice_1 := arr[:i-1]//10
	slice_2 := arr[i:]//30,40,50

	new_slice := append(slice_1,slice_2...)

	fmt.Println(new_slice)
}

//COPYING FROM ONE SLICE TO ANOTHER(Note that when you change a value in the dest slice, the src slice does not get affected )
// func main() {
// 	src_slice := []int{10, 20, 30, 40, 50}
// 	dest_slice := make([]int, 3)

// 	num := copy(dest_slice, src_slice)

// 	dest_slice[1] = 900

// 	fmt.Println(src_slice)
// 	fmt.Println(dest_slice)
// 	fmt.Println(num)
// }

EXERCISES

func main(){
	arr := [5]int{}
	my_map := make(map[string]int)
	my_map["A"] = 65
	my_map["B"] = 66
	i := 0
	for _, value := range my_map {
			arr[i] = value
			i += 1
	}
	fmt.Println(arr)
}

func main() {
	arr := [5]int{10, 20, 30, 90, 100}
	new_slice := append(arr[:3], arr[4:]...)
	fmt.Print(new_slice)
}

MAPS

NOTE THAT THIS CREATES A NIL MAP AND WE CANNOT ADD VALUES TO IT
func main(){
	var my_map map[string]int
	fmt.Println(my_map)
}

TO ADD PAIRS TO THE MAP , WE NEED TO CREATE AND INITIALIZE THE MAP

func main(){
	code := map[string]string{"en":"English", "fr":"French"}
	fmt.Println(code)
}

WE CAN ALSO INITIALIZE A MAP USING THE MAKE FUNCTION (this gives an empty map)

func main(){
	code := make(map[string]int)
	fmt.Println(code)
}

GETTING A VALUE FROM A MAP(When you get a value, it returns a value and a boolean that shows if it was found or not)

func main(){
	code := map[string]int{"en":1, "fr":2}
	value, found := code["en"]
	fmt.Println(value, found)
}

ADDING A KEY VALUE PAIR

func main(){
	code := map[string]int{"en":1, "fr":2}
	code["en"] = 200
	value, found := code["en"]
	fmt.Println(value, found)
}

DELETING A VALUE

func main(){
	code := map[string]int{"en":1, "fr":2}
	code["en"] = 200
	delete(code, "en")
	value, found := code["en"]
	fmt.Println(value, found)
}

ITERATING OVER A MAP
func main(){
	code := map[string]string{"En":"English", "Fr":"French"}

	for key, value := range code {
		fmt.Println(key, "=>", value)
	}
}