package main

import (
	"fmt"
	"math"
)

// interface definition
// type VowelsFinder interface {
// 	FindVowels() []rune
// }

// //type Mystring is created
// type MyString string

// //MyString implements VowelsFinder
// func (ms MyString) FindVowels() []rune {
// 	//create a run variable
// 	var vowels []rune
// 	//iterate over ms and print out the rune
// 	for _, rune := range ms {
// 		if rune == 'a' || rune == 'e' || rune == 'o' || rune == 'u' {
// 			vowels = append(vowels, rune)
// 		}
// 	}
// 	return vowels
// }

// func main(){
// 	name := MyString("Sam Anderson")
// 	var v VowelsFinder
// 	//we  assign name which is of type MyString to v of type VowelsFinder
// 	v = name //possible since MyString implements VowelsFinder
// 	fmt.Printf("Vowels are %c", v.FindVowels())
// }

// //EXAMPLE 2
// //Program that calculates the total expense for a company based on the individual salaries of the employees

// type SalaryCalculator interface {
// 	CalculateSalary() int
// }

// //the two kinds of employees in the company are defined by structs
// type Permanent struct {
// 	empId  		int
// 	basicpay 	int
// 	pf			int
// }

// type Contract struct {
// 	empId		int
// 	basicpay	int
// }

// type Freelancer struct {
// 	empId		int
// 	ratePerHour	int
// 	totalHours	int
// }

// //here both Permanent and Contract structs implement the Salary calculator interface

// //salary or permanent employee is the sum of basic pay and pf
// func (p Permanent) CalculateSalary() int {
// 	return p.basicpay + p.pf
// }

// //salary of contract employee is the basic pay alone
// func (c Contract) CalculateSalary() int {
// 	return c.basicpay
// }

// //salary of freelancer
// func (f Freelancer) CalculateSalary() int {
// 	return f.ratePerHour * f.totalHours
// }

// /*
// total expense is calculated by iterating through the salries
// of the individual employees
// */

// //this function takes in a slice of the salarycalculator interface
// //the totalExpense function is utilizing the interfaces above to calculate the total expense of both the permanent and temporary employees

// func totalExpense(s []SalaryCalculator) {
// 	expense := 0
// 	//for each employee (v)
// 	//we are able to call v.calculate salary for the respective employees because they implement the interface which has the calculatesalary method
// 	for _, v := range s {
// 		expense = expense + v.CalculateSalary()
// 	}
// 	fmt.Printf("Total Expense Per Month $%d", expense)
// }

// func main(){
// 	pemp1 := Permanent{
// 		empId: 1,
// 		basicpay: 5000,
// 		pf: 20,
// 	}

// 	pemp2 := Permanent{
// 		empId: 2,
// 		basicpay: 6000,
// 		pf: 30,
// 	}

// 	cemp1 := Contract{
// 		empId: 3,
// 		basicpay: 3000,
// 	}

// 	femp1 := Freelancer{
// 		empId: 4,
// 		ratePerHour: 70,
// 		totalHours: 120,
// 	}

// 	//here we create the Salary interface slice that contains the permanent and Contract and freelance types
// 	//we are able to this because these types implement the salarycalculator interface
// 	employees := []SalaryCalculator{pemp1, pemp2, cemp1, femp1}

// 	//the totalExpense function literally iterates through this slice as defined above
// 	totalExpense(employees)
// }

// //we can basically loop through the employees and implement calculateSalary for them respectively
// //they each implement their own version of implementSalary

// //IMPORTANT
// //it is important to note that , when you create a variable for the interface, the variable could be
// // a normal variable for the interface , like the example we have above (v) or could be a slice of the interface([]Salary Calculator)

// //now this variable is assigned to the type that implemented it. In the first example, it was assigned to name of type Mystring
// //and in the second example it was assigned to type permamnent and company which both implemented it. but this time
// //it was assigned as a slice because we have multiple of it

// //The biggest advantage of this is that totalExpense can be extended to any new employee without any code changes
// //Add a new type of employee Freelancer with a differnet salary structure

// // SIMPLE STEPS TO IMPLEMENTING INTERFACE
// //1. create the interface
// //2. create the type that will implement the interface(e.g struct)
// //3. the struct defines the interface method(i.e a method for what the interface actually does for the struct)e.g func(c Circle) Area() float64
// //4. in the main function, create instances of the struct
// //5. in the main function, also create a variable for the interface and assign it to the struct(note that you can pass in the struct directly to the function)
// //6. In the main function, call the function and pass in that variable
// //7. the function in step 6 takes in the interface variable(this could be single(example1) or a slice of variables(example 2)) and recall this variable stores the struct type that is implementing it
// //8. inside that function we can now do struct_instance.method() e.g v.CalculateSalary in example 2
// // 9. note that a function does not always have to be created like in example 2, step 7 and 8 can also be called in the main function like in example 1

// //EXAMPLE 3
// //An interface can be thought of as being represented internally by a tuple(type, value)
// //type is the underlying concrete type of the interface and value holds the value of the concrete type

// type Worker interface {
// 	Work()
// }

// type Person struct {
// 	name string
// }

// func (p Person) Work() {
// 	fmt.Println(p.name, "is working")
// }

// func describe(w Worker){
// 	fmt.Printf("Interface type %T value %v\n", w, w)

// }

// func main(){
// 	p := Person{name: "Naveen"}

// 	var w Worker = p

// 	describe(w)
// 	w.Work()
// 	//here we called the work method in the main function unlike what we did in example 2

// }

// //USING THE STRINGER INTERFACE//EXAMPLE 4

// type Article struct {
// 	Title 	string
// 	Author 	string
// }

// type Book struct {
// 	Title 	string
// 	Author 	string
// 	Pages 	int
// }

// type Stringer interface {
// 	String() string
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
// }

// func (b Book) String() string {
// 	return fmt.Sprintf("The %q book was written by the Author %s", b.Title,b.Author)
// }
// func main() {
// 	a := Article {
// 		Title: "Understanding Interfaces in Go",
// 		Author: "Sammy Shark",
// 	}
// 	//this is an example where we did not have to assign the struct to the interface variable first(we pass the struct variable to the fucntion directly)
// 	Print(a)

// 	//we can add another struct to use the stringer interface
// 	b := Book {
// 		Title: "Learning about GO",
// 		Author: "Alero Awani",
// 		Pages: 34,
// 	}

// 	Print(b)

// }

// func Print(s Stringer) {
// 	fmt.Println(s.String())
// }

// //here the Print method takes a Stringer and not a concrete type of Article
// //because the compiler knows that a Stringer interface defines the string method, it will only accept types that also have the string method
// //hence we can use the Print method with anything that satisfies the stringer interface

// //EXAMPLE 5

// //EXAMPLE WITH CONCRETE TYPE// this is an example of a the print method taking in a concrete type and not the Stringer interface//Not a good practice
// type Article struct {
// 	Title string
// 	Author string
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
// }

// func main() {
// 	a := Article {
// 		Title: "Understanding Interfaces in Go",
// 		Author: "Sammy Shark",
// 	}

// 	Print(a)

// }

// func Print(a Article) {
// 	fmt.Println(a.String())
// }// with this concrete type, you can see that the interface was not defined here. and Print takes in the concrete type(a Article) this limits the
// //method to only a

// //THERE ARE TWO WAYS OF IMPLEMENTING AN INTERFACE
// //1. As an argument to a function - best practice
// //2. As a concrete type

//EXAMPLE 6
//Define the struct
type Circle struct{
	Radius float64
}

type Square struct {
	Width float64
	Height float64
}

//define the interface
type Sizer interface {
	Area() float64
}


//this interface is a combination of two 
//we didnt have to create the stringer interface again because it is already part of the Go fmt package
type Shaper interface {
	Sizer
	fmt.Stringer
}

//the both structs will define the method Area() of the interface Sizer
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

//add the string() method as a behaviour to each of the struct types(this will satisfy the fmt.Stringer interface)
func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, height: %.2f}", s.Width, s.Height)
}

//call the Area method from the square and circle types in a function
// This function can be used by anything that satisfies the Sizer interface

//here we take both arguments as the type Sizer and return the result
//as a Sizer as well. This means that we no longer return a Square or a Circle, but the interface of Sizer
func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

//use the Shaper interface in a function which calls both the String() method and Area method

func PrintArea(s Shaper) {
	fmt.Printf("Area of %s is %.2f\n", s.String(), s.Area())
}




//Define the main function 
func main(){
	//create instances of structs 
	c := Circle{Radius: 10}
	PrintArea(c)

	s := Square{Height: 10, Width: 5}
	PrintArea(s)

	l := Less(c, s)
	fmt.Printf("%+v is the smallest\n",l)


}











//Another way to look at is, if more than one type(structs) is implementing the same method you can create an interface for it instead 
