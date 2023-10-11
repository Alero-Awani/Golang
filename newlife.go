package main

// import "fmt"

// func main(){
// 	var name string
// 	var is_muggle bool

// 	fmt.Print("Enter your name and are you a muggle? ")
// 	count, err := fmt.Scanf("%s %t", &name, &is_muggle)
// 	fmt.Scanf("%s %t", &name, &is_muggle)
// 	fmt.Println(name, is_muggle)
// 	fmt.Println("Count: ", count)
// 	fmt.Println("err: ", err)
// }

// func main(){
// 	vals := make([]int, 5)

// 	fmt.Println("Capacity was: ", cap(vals))

// 	for i := 0; i < 5; i++{
// 		vals = append(vals, i)
// 	}
// 	fmt.Println("Capacity is now", cap(vals))

// }

// func addNumbers(a int, b int)int{
// 	sum := a + b

// 	return sum
// }

// func main(){
// 	sumNumbers := addNumbers(10,20)

// 	fmt.Print(sumNumbers)
// }

// type Circle struct {
// 	x float64
// 	y float64
// 	r float64
// }

// type Student struct {
// 	name string
// 	rollNo int
// 	marks []int
// 	grades map[string]int
// }

// func main(){
// 	st := Student {
// 		name: "Joe",
// 		rollNo: 12,
// 	}
// 	fmt.Printf("%+v", st)
// }

//FIND THE FREQUENCY OF WORDS IN A SENTENCE
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func wordFrequency(text string) map[string]int {
// 	my_map := make(map[string]int)
// 	words := strings.Fields(text)
// 	for _, value := range words {
// 		my_map[value]++
// 	}
// 	return my_map
// }

// func main() {
// 	text := "The quick brown fox jumps over the lazy dog"
// 	fmt.Println(wordFrequency(text))

// }

//  we need to implement a function that will return a list of items that are in a given price range.
//below what gets passed into the function is a array of struct
// package main

// import "fmt"

// type Item struct {
// 	Name  string
// 	Price float64
// }

// Not that the function below is not a method
// func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {
// 	// Initialize an empty slice called "result" to store the items that match the price range.
// 	var result []Item

// 	// Iterate over the "items" slice using a for loop.
// 	for _, item := range items {
// 		if item.Price >= minPrice && item.Price <= maxPrice {
// 			// Iterate over the "items" slice using a for loop.
// 			result = append(result, item)
// 		}
// 	}
// 	// Return the "result" slice containing items that match the price range.
// 	return result

// }

// func main() {
// 	// here we are initializing the array of structs
// 	items := []Item{
// 		{Name: "Apple", Price: 0.5},
// 		{Name: "Banana", Price: 0.25},
// 		{Name: "Orange", Price: 0.75},
// 		{Name: "Pineapple", Price: 1.5},
// 	}

// 	fmt.Println(getItemsInPriceRange(items, 0.0, 1.0))
// 	fmt.Println(getItemsInPriceRange(items, 0.5, 1.0))
// 	fmt.Println(getItemsInPriceRange(items, 0.75, 1.5))
// }

// We are building a program that helps users track their expenses. We need to create a struct called Expense to store information about an individual expense, including the name of the expense, the amount, and the date.
// We need to create a method called Total that calculates the total amount spent on expenses.
// Also, we need to create a method called getName on Expense struct that returns the name of the Expense.

// package main

// import "fmt"

// Declare the Expense struct here
// type Expense struct {
// 	expense string
// 	amount float64
// 	date string

// }

// // Implement the Total method to calculate the total amount spent
// func Total(e []Expense) float64 {

// 	var total float64 = 0

// 	for _, value := range e {
// 		total += value.amount
// 	}

// 	return total

// }
// // Implement the get Name method on the Expense struct here  - here we want to
// func (e *Expense) getName() string{
// 	return e.expense
// }

// // here we have a list of struct
// func main() {
// 	expenses := []Expense{
// 		{"Grocery", 50.0, "2022-01-01"},
// 		{"Gas", 30.0, "2022-01-02"},
// 		{"Restaurant", 40.0, "2022-01-03"},
// 	}

// 	fmt.Println(Total(expenses))
// 	fmt.Println(expenses[0].getName())
// }