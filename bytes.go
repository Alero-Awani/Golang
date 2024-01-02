package main

import "fmt"

//https://syslog.ravelin.com/byte-vs-string-in-go-d645b67ca7ff#:~:text=You%20can%20change%20the%20bytes,creating%20a%20whole%20new%20string.

//A string is a sequence of bytes

// A byte is 8 bits(1s and 0s)
/*
[]byte and string are small headers pointing to data,
 with lengths indicating how much data is present.
[]byte has two lengths: the CURRENT LENGTH of the data and
the CAPACITY. The capacity tells us how many more bytes we can add to the
 data before needing to go and get a bigger piece of memory.
*/

/*
[]byte is not immutable. You can change the bytes within a slice.
 You can grow the slice by appending bytes to the end. But a string
 is immutable. You cannot change it and you cannot grow it without creating
 a whole new string.
*/

/*
[]byte
type slice struct {
    data uintptr
    len int
    cap int
}

string
type string struct {
    data uintptr
    len int
}
*/

//string literals in Golang are encoded in UTF-8

//CODE POINT - In utf-8 a code point is a number that represents the bytes(hex values) of a character
//it could map to between 1 to 3 bytes. for the chinese character that has 3 bytes. One code point numeral maps to those code points

//for letter H, the hex numebr and the hex code point are the same
//this happens for a lot of english letters that usually have one byte

func main(){
	var a = "Hello there"//string literal encoded at UTF-8
	var b = "六书"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("H length:", len("H")) //1 //this returns the number of bytes that represents the character
	fmt.Println("六 length: ", len("六")) //3 bytes

	//print out the hexadecimal , decimal and binary for the one byte in H
	for i := 0; i < len("H"); i++ {
		fmt.Printf("H hexadecimal: %x decimal: %3d binary: %o8b", "H"[i],"H"[i],"H"[i])
	}

	fmt.Print("\n")

	//using our outputs we can print out the letter H

	fmt.Println(string(0x0048)) //using hexadecimal(hex number is formatted for the compiler to understand)
	fmt.Println(string(72))// using decimal


	//looping through the character with 3 bytes 

	for i := 0; i < len("六"); i++ {
		fmt.Printf("六 hexadecimal: %x decimal: %3d binary: %o8b\n", "六"[i],"六"[i],"六"[i])
	}

	//using the hex code point to print out the value 
	fmt.Println(string(0x516D))

	//using the decimal code point to print out the character 
	fmt.Println(string(20845))

	//the first three indexes which represent the 3 bytes from the variable b will print out the first character 

	fmt.Println(string(b[0:3]))
	fmt.Println(string(b[3:6]))

	//now lets print out the length of both variables 
	fmt.Println("length of Hello there:", len("Hello there"))

	fmt.Println("Length of 六书:", len("六书"))//both these chars are represented by three bytes hence the len is 6


	//loop through the characters in hello there and use the hexadecimal values to print the word again

	for index, s := range "Hello there" {
		fmt.Printf("%3c is index number %3d decimal: %6d hexadecimal: %x binary: %08b\n", s, index, s, s, s)
	}

	//using the hexadecimal values , we can print hello as such 
	fmt.Print(string(0x0048), string(0x0065), string(0x006c), string(0x006c), string(0x006f), "\n")


	//do the same for the chinese characters
	for index, s := range "六书" {
		fmt.Printf("%3c is index number %3d decimal: %6d hexadecimal: %x binary: %08b\n", s, index, s, s, s)
	}

	fmt.Print(string(0x516d), string(0x4e66), "\n")


	//recall that a string is a slice of bytes
	//lets create our own slice of bytes 

	byteSlice1 := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}// 0x0048 works too

	fmt.Println("string(byteSlice1):", string(byteSlice1))

	//here we are trying to create the byte slice for the chinese characters
	//but it doesnt work because the hexadecimal number is too big to be represented in a single byte 
	// byteSlice2 := []byte{0x516d, 0x4e66}

	//instead we can use a slice of runes 
	//slice is uint8 while rune is int 32 so it can actually hold up to 4 bytes and the ones for the chines characters are 3 bytes each

	runeSlice1 := []rune{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	fmt.Println("string(runeSlice1):", string(runeSlice1))

	runeSlice2 := []rune{0x516d, 0x4e66}
	fmt.Println("string(runeSlice2):", string(runeSlice2))



}


//you can get the hex code point and Decimal code point for characters using this tool 
//https://www.cogsci.ed.ac.uk/~richard/utf-8.html


//when you see int32 thats a clue that youre working with runes 