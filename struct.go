package main

// import "fmt"

// type Student struct {
// 	name string
// 	rollNo int
// 	marks []int
// 	grades map[string]int
// }

// type Student struct {
// 	name string
// 	rollNo int
// 	marks []int
// 	grades map[string]int
// }

// func main(){
// 	st := Student{"Joe", 12}
// 	fmt.Printf("%+v",st)
// }

//INITIALIZING A STRUCT
//ONE
// var s Student

//TWO - USING THE NEW KEYWORD
// func main(){
// 	st := new(Student)
// 	fmt.Printf("%+v", st)
// }

//THREE - USING SHORT VARIABLE ASSIGNMENT
// func main(){
// 	st := Student{
// 		name: "Joe",
// 		rollNo: 12,
// 	}
// 	fmt.Printf("%+v", st)

// }

//OR - OMIT FIELD NAMES AND SPECIFY THE VALUES IN THE ORDER THEY ARE DECLARED- not recommended for maintainability
// func main(){
// 	st := Student{"Joe", 12}
// 	fmt.Printf("%+v", st)
// }

//ACCESSING FIELDS OF A STRUCT AND MODIFYING THEM
//When we want to change the value of the field of a structure, we use the passing by ref method

// type Circle struct {
// 	x int
// 	y int
// 	radius float64
// 	area float64
// }

// func calcArea(c *Circle){
// 	const PI float64 = 3.14
// 	var area float64
// 	area = (PI * c.radius * c.radius)
// 	(*c).area = area
// }

// func main(){
// 	c := Circle{x:5, y:5, radius:5, area:0}
// 	fmt.Printf("%+v \n", c)
// 	calcArea(&c)
// 	fmt.Printf("%+v \n", c)
// }

// type Circle struct {
// 	x int
// 	y int
// 	radius float64
// 	area float64
// }

// func calcArea(c *Circle){
// 	const PI float64 = 3.14
// 	var area float64
// 	area = (PI * c.radius * c.radius)
// 	(*c).area = area

// }

// func main(){
// 	c := Circle{
// 		x:5,
// 		y:5,
// 		radius:5,
// 		area:0,
// 	}
// 	fmt.Printf("%+v \n", c)
// 	calcArea(&c)
// 	fmt.Printf("%+v \n", c)
// }

// EXAMPLE 1 - This returns an error because mov is supposed to be passed in as &mov

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 			name:   s,
// 			rating: r,
// 	}
// 	return
// }

// func increaseRating(m *Movie) {
// 	m.rating += 1.0
// }

// func main() {
// 	mov := getMovie("xyz", 2.0)
// 	increaseRating(mov)
// 	fmt.Printf("%+v", mov)
// }

//EXAMPLE 2

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 			name:   s,
// 			rating: r,
// 	}
// 	return
// }

// func main() {
// 	mov := getMovie("xyz", 2.1)
// 	mov1 := getMovie("abc", 3.3)
// 	fmt.Println(mov)
// 	fmt.Printf("%+v \n", mov)
// 	movies := make([]Movie, 5)
// 	fmt.Println(movies)
// 	movies = append(movies, mov)
// 	movies = append(movies, mov1)
// 	fmt.Println(movies)
// 	// for _, value := range movies {
// 	// 		fmt.Println(value)
// 	// }
// }

// type Movie struct {
// 	name string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 		name: s,
// 		rating: r,
// 	}
// 	return
// }

// func main() {
// 	mov := getMovie("xyz", 2.1)
// 	mov1 := getMovie("abc", 3.3)

// 	movies := make([]Movie, 5)
// 	fmt.Println(movies)

// 	movies = append(movies, mov)
// 	movies = append(movies, mov1)

// 	fmt.Println(movies)

// 	for _, value := range movies {
// 		fmt.Println(value)
// 	}
// }

// EXAMPLE 3

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func main() {
// 	mov := Movie{"xyz", 2.1}
// 	mov1 := Movie{"abc", 2.1}
// 	if mov.rating == mov1.rating || mov != mov1 {
// 			fmt.Println("condition met")
// 	} else if mov.rating == mov1.rating {
// 			fmt.Println("condition_2 met")
// 	}
// }

//COMAPARING STRUCT
// type s1 struct {
// 	x int
// }

// func main(){
// 	c := s1{x:5}
// 	c1 := s1{x:6}
// 	c2 := s1{x:5}

// 	if c != c1 {
// 		fmt.Println("c and c1 have different values")
// 	}
// 	if c == c2 {
// 		fmt.Println("C is the same as C2")
// 	}
// }