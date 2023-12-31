package main

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
// )

// //allow the user specify what text message they want printed on the console by using a flag
// //if the user doesnt specify a flag then we want a default message to be printed

// func main(){
// 	//set  a flag to specify the message to display
// 	//text is the name of the flag, what comes after is the default message , then a help message to show how the flag works
// 	// not that the flag returns a pointer, so msg is a pointer
// 	//recall that if you want to use a pointer you need to dereference it
// 	msg := flag.String("text", "Hello there!","message to display")

// 	caps := flag.Bool("c", false, "should text be uppercase")
// 	nums := flag.Int("n", 1, "number of times text is displayed")
// 	flag.Parse()

// 	if *caps {
// 		*msg = strings.ToUpper(*msg)
// 	}

// 	for i := 0; i < *nums; i++ {
// 		fmt.Println(*msg)
// 	}

// }

// //run go run flag.go -text="Goodmorning"