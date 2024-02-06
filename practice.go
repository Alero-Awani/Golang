package main

// import "fmt"

// func main(){
// 	var user, country string
// 	user = "Alero"
// 	country = "Nigeria"
// 	fmt.Println(user, country)
// }

// func main(){
// 	var(
// 		s string = "foo"
// 		i int = 5
// 	)

// 	fmt.Println(s, i)
// }

// func main(){
// 	name := "Lisa"
// 	name = "Alero"

// 	fmt.Println(name)
// }

// func main(){
// 	var name string
// 	fmt.Print("Enter your name here: ")
// 	fmt.Scanf("%s", &name)
// 	fmt.Println("Hey there", name)
// }

// func main(){
// 	var grades int = 42
// 	fmt.Printf("The value of the grade is %v and it is of type %T", grades, grades)
// }

// func main(){
// 	var i int = 43
// 	var f float64 = float64(i)
// 	fmt.Printf("%.2f \n", f)

// }

// func main(){
// 	var f float64 = 45.87
// 	var i int = int(f)
// 	fmt.Printf("%v \n", i)
// }

// Convert string to Integer

// func main(){
// 	var s string = "200"
// 	i, err := strconv.Atoi(s)
// 	fmt.Printf("%v, %T \n", i, i)
// 	fmt.Printf("%v, %T", err, err)
// }

// const PI float64 = 3.14

// func main(){
// 	var radius float64 = 43.2
// 	var area float64

// 	area = PI * radius * radius

// 	fmt.Printf("Radius: %.2f \nPI: %.1f\n", radius, PI)
// 	fmt.Printf("The areas is: %f \n", area)
// 	fmt.Println("Thank you!")
// }

//COMPARISON OPERATORS
// func main(){
// 	var city string = "kolkata"
// 	var city_2 string = "canada"

// 	fmt.Println(city == city_2)
// 	fmt.Println(city != city_2)

// }

// func main(){
// 	for i := 1; i <= 5; i++ {
// 		if i == 3 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// }

// func main(){
// 	var grades [5]int = [5]int{90, 80, 70}
// 	grades[1] = 100
// 	fmt.Println(grades)
// }

// func main(){
// 	var grades [5]int = [5]int{90, 80, 70}
// 	for i := 0; i < len(grades); i++ {
// 		fmt.Println(grades[i])
// 	}
// }

// func main(){
// 	var grades [5]int = [5]int{90,80,70}
// 	for index, element := range grades {
// 		fmt.Println(index, "=>", element)
// 	}
// }

// MULTIDIMENSIONAL ARRAYS

// func main(){
// 	arr := [3][2]int{{2,4},{3,4}}
// 	for index, element := range arr{
// 		fmt.Println(index, "=>", element)
// 	}
// }

// func main(){
// 	slice := []int{10,20,30}
// 	fmt.Println(slice)
// }

// func main(){
// 	slice := make([]int,5,8)
// 	fmt.Println(slice)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))
// }

// func main(){
// 	arr := [5]int{10,20,30,40,50}
// 	slice := arr[1:3]

// 	slice[0] = 25

// 	fmt.Println(arr)
// 	fmt.Println(slice)
// }

// func main(){
// 	arr := [5]int{10,20,30,40,50}
// 	i := 2
// 	fmt.Println(arr)
// 	slice_1 := arr[:i]
// 	slice_2 := arr[i+1:]

// 	new_slice := append(slice_1, slice_2...)

// 	fmt.Println(new_slice)
// }

// func main(){
// 	code := map[string]string{"en":"English", "jp":"Japan"}
// 	value, _ := code["en"]
// 	delete(code, "jp")
// 	fmt.Println(value)
// 	fmt.Println(len(code))
// }

// func main(){
// 	arr := [5]int{}
// 	my_map := make(map[string]int)
// 	my_map["A"] = 65
// 	my_map["B"] = 66
// 	i := 0
// 	for _, value := range my_map {
// 			arr[i] = value
// 			i += 1
// 	}
// 	fmt.Println(arr)
// }

// func main() {
// 	arr := [5]int{10, 20, 30, 90, 100}
// 	new_slice := append(arr[:3], arr[4:]...)
// 	fmt.Print(new_slice)
// }

// func main(){
// 	arr := [5]int{}
// 	fmt.Println(arr)

// 	my_map := make(map[string]int)
// 	my_map["A"] = 65
// 	my_map["B"] = 66

// 	fmt.Println(my_map)
// 	i := 0

// 	for _, value := range my_map {
// 		arr[i] = value
// 		i += 1
// 	}

// 	fmt.Println(arr)
// }

// func operation(a int, b int)(int, int){
// 	sum := a + b
// 	diff := a - b

// 	return sum, diff
// }

// func main(){
// 	sum, diff := operation(1, 2)

// 	fmt.Println(sum, diff)
// }

//CALCULATING THE AREA/PERIMETER ETC of a circle without using HOF
// func calcArea(r float64) float64{
// 	return 3.14 * r * r
// }

// func calcPerimeter(r float64) float64{
// 	return 2 * 3.14 * r
// }

// func calcDiamater(r float64) float64{
// 	return 2 * r
// }

// func main(){
// 	var query int
// 	var radius float64

// 	fmt.Printf("Enter the radius of the circle: ")
// 	fmt.Scanf("%f", &radius)
// 	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
// 	fmt.Scanf("%d", &query)

// 	if query == 1 {
// 		fmt.Println("Result: ", calcArea(radius))
// 	} else if query == 2 {
// 		fmt.Println("Result: ", calcPerimeter(radius))
// 	} else if query == 3 {
// 		fmt.Println("Result: ", calcDiamater(radius))
// 	} else {
// 		fmt.Println("Invalid query")
// 	}

// }

// USING HIGHER ORDER FUNCTIONS TO CARRY OUT THE SAME THING
// func calcArea(r float64) float64{
// 	return 3.14 * r * r
// }

// func calcPerimeter(r float64) float64{
// 	return 2 * 3.14 * r
// }

// func calcDiameter(r float64) float64{
// 	return 2 * r
// }

// func printResults(radius float64, calcFunction func(r float64)float64) {
// 	result := calcFunction(radius)
// 	fmt.Println("Result: ", result)
// 	fmt.Println("Thank you")
// }

// func getFunction(query int) func(r float64)float64{
// 	query_to_func := map[int]func(r float64)float64 {
// 		1: calcArea,
// 		2: calcPerimeter,
// 		3: calcDiameter,
// 	}
// 	return query_to_func[query]
// }

// func main(){
// 	var query int
// 	var radius float64

// 	fmt.Printf("Enter the radius of the circle: ")
// 	fmt.Scanf("%f", &radius)
// 	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
// 	fmt.Scanf("%d", &query)

// 	printResults(radius, getFunction(query))
// }
// // here you can see the function is taking in another function as an input
// func printResults(radius float64, calcFunction func(r float64)float64){
// 	//store the output of the function in variable result
// 	result := calcFunction(radius)
// 	fmt.Println("Result: ",result)
// 	fmt.Println("Thank you")
// }

// // this function takes in an input of query and returns a function that takes in a float input and returns a float output
// func getFunction(query int) func(r float64)float64{
// 	// this function consists of a map that maps query numbers to functions
// 	// we can see when creating the map that the key data type is int and the value data type is func
// 	query_to_func := map[int]func(r float64)float64{
// 		1: calcArea,
// 		2: calcPerimeter,
// 		3: calcDiameter,
// 	}
// 	return query_to_func[query]
// }

// func main(){
// 	var query int
// 	var radius float64

// 	fmt.Printf("Enter the radius of the circle: ")
// 	fmt.Scanf("%f", &radius)
// 	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
// 	fmt.Scanf("%d", &query)

// 	printResults(radius, getFunction(query))

// }

// // EXAMPLE
// //here we have a function that takes in an integer and returns an integer, the func only adds 100 to the number
// func addHundred(x int) int {
// 	return x + 100
// }

// //here we have a func that takes in any amount of int and returns an anonymous func
// func partialSum(x ...int) func() {
// 	sum := 0
// 	// incrementing the sum by whatever range of values is put in
// 	for _, value := range x {
// 			sum += value
// 	}

// 	return func() {
// 			fmt.Println(addHundred(sum))
// 	}
// }
// func main() {
// 	partial := partialSum(1, 2, 3, 4, 5)
// 	partial()
// }

// func addHundred(x int)int {
// 	return x + 100
// }

// func partialSum(x ...int) func() {
// 	sum := 0
// 	for _, value := range x {
// 		sum += value
// 	}
// 	return func() {
// 		fmt.Println(addHundred(sum))
// 	}
// }

// func main(){
// 	partial := partialSum(1,2,3,4,5,6)
// 	partial()
// }

// EXAMPLE 2 - THIS RETURNS NO OUTPUT, this is because a function must have a return statement if the output parameter type was specified unlike the function we
//have above that didnt have it specified

// func addHundred(x int) int {
//         return x + 100
// }
// func partialSum(x ...int) func() int {
//         sum := 0
//         for _, value := range x {
//                 sum += value
//         }
// 		// this returns an integer and not a function
//         return func() int {
// 			// fmt.Println(addHundred(sum))
// 			return addHundred(sum)
//         }
// }
// func main() {
//         partial := partialSum(1, 2, 3)
//         partial()
// }

//EXAMPLE 3

// func addHundred(x int) int {
//         return x + 100
// }

// //here add100 is the name of the function variable, when calling the function we pass in addHundred=> answer is 106
// func partialSum(add100 func(x int) int, x ...int) int {
//         sum := 0
//         for _, value := range x {
//                 sum += value
//         }
//         return add100(sum)

// }
// func main() {
//         partial := partialSum(addHundred, 1, 2, 3)
//         fmt.Println(partial)
// }

//EXAMPLE 4
// func addHundred(x int) {
// 	fmt.Println(x + 100)
// }
// func partialSum(add100 func(x int), x ...int) int {
// 	sum := 0
// 	for _, value := range x {
// 			sum += value
// 	}
// 	add100(sum)
// 	return 0
// }
// func main() {
// 	partial := partialSum(addHundred, 1, 2, 3)
// 	fmt.Println(partial)
// }

// func main(){
// 	i := 10
// 	fmt.Printf("%T %v \n", &i, &i)
// 	fmt.Printf("%T %v \n", *(&i), *(&i))
// }

//DECLARING AND INITIALIZING A POINTER USING THREE METHODS
//ONE

// func main(){
// 	i := 10
// 	var ptr_i *int = &i
// 	fmt.Println(ptr_i)
// }

//TWO - HERE WE DO NOT SPECIFY THE DATA TYPE BUT IT WILL BE INTERNALLY DETERMINED BY THE COMPILER

// func main(){
// 	s := "hello"
// 	var ptr_s = &s
// 	fmt.Println(ptr_s)
// }

// THREE - USING THE SHORTHAND OPERATOR

// func main(){
// 	s := "hello"
// 	ptr_s := &s
// 	fmt.Println(ptr_s)
// }

//USING THE DEREFERNCING OPERATOR WITH POINTERS

// func main(){
// 	s := "hello"
// 	fmt.Printf("%T %v \n", s, s)

// 	ps := &s
// 	fmt.Printf("%T \n", &s)
// 	*ps = "world"

// 	fmt.Printf("%T %v \n",s,s)
// }

// func main(){
// 	s := "hello"
// 	fmt.Printf("%T %v \n", s, s)

// 	ps := &s
// 	fmt.Printf("%T \n", &s)
// 	*ps = "world"

// 	fmt.Printf("%T %v \n", s, s)
// }


