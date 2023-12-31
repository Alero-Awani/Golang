package main

// import (
// 	"fmt"
// 	"sync"
// )

// // func someFunc(num string) {
// // 	fmt.Println(num)
// // }

// // func main() {
// // 	time.Sleep(2)
// // 	go someFunc("1")
// // 	fmt.Println("Hi there, this is from the main function")
// // }

// // func main(){
// // 	myChannel := make(chan string)

// // 	go func(){
// // 		myChannel <- "data"
// // 	}()

// // 	msg := <- myChannel

// // 	fmt.Println(msg)
// // }


// // here we encounter a deadlock because the main function consumes only 4 of the data from the channel when 5 was sent 
// func countCat(c chan string, wg *sync.WaitGroup){
// 	for i := 0; i < 5; i++{
// 		c <- "Cat"
// 	}
// 	wg.Done()
// }


// func main(){
// 	myChannel := make(chan string)
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)

// 	go countCat(myChannel, &wg)

// 	for i := 0; i < 4; i++{
// 		msg := <- myChannel
// 		fmt.Println(msg)
// 	}
// 	wg.Wait()

// }