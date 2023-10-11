package main

// import "fmt"

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


//USING HIGHER ORDER FUNCTIONS TO CARRY OUT THE SAME THING 
// func calcArea(r float64) float64{
// 	return 3.14 * r * r
// }

// func calcPerimeter(r float64) float64{
// 	return 2 * 3.14 * r 
// }

// func calcDiameter(r float64) float64{
// 	return 2 * r
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
// //here we have a function that takes in an integer and returns an integer, the fun only adds 100 to the number
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