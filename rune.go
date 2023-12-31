package main

//code formating for rune is with %c

//rune is a built in type in go and is an alias of int32. it represents a unicode code point in go
//A unicode code point is a (numerical value) that uniquely identifies a character in the unicode standard
//runes are essential for working with non-ASCII characters and multilingual text

//unlike bytes which are used to stroe ASCII characters, runes can represent characters from various languages and symbol sets

// func main(){
// 	//Decalre a rune variable (runes are enclosed in single quote)
// 	var r rune = 'A'

// 	//print the run and its corresponding code point
// 	fmt.Printf("Rune: %c\nCode Point: %U\n", r, r)
// }

//runes are important because they allow us to work with characters from different languages and symbol sets.
//here we iterate over a japanese string and print each rune

// func main(){
// 	str := "こんにちは, 世界!"
// 	for _, r := range str {
// 		fmt.Printf("%c", r)
// 	}
// }

// func printBytes(s string) {
// 	fmt.Printf("Bytes: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func printChars(s string) {
// 	fmt.Printf("Characters: ")
// 	// here the string is converted to a slice of runes
// 	runes := []rune(s)
// 	// we loop over the slice and print the characters
// 	for i := 0; i < len(runes); i++ {
// 		fmt.Printf("%c ", runes[i])
// 	}
// }

// func main() {
// 	name := "Hello World"

// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)

// 	fmt.Printf("\n")

// 	printBytes(name)
// }