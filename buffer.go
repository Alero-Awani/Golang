package main

import (
	"bytes"
	"fmt"
)

//buffer belongs to the byte package
//if a string is too long for example and we want to calculate its length
//instead of using len we can use the buffer to calculate it in the form of chunks of data

//buffer allows us to give buffer storage where we can store some data
//and access the same data when needed

//Often we want to build up a long sequence of bytes. With bytes.Buffer
//we can write bytes into a single buffer, and then convert to a string when we are done.


func main(){
	//New Buffer
	var b bytes.Buffer


	//Write strings to the buffer
	b.WriteString("ABC")
	b.WriteString("DEF")


	//Convert to a string and print it
	fmt.Println(b.String())
}