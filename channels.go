package main

// import (
// 	"fmt"
// )

// func main() {
// 	// create two go routine and use a channel to communicate between them
// 	ch := make(chan string)

// 	go sell(ch)
// 	go buy(ch)
// 	time.Sleep(2 * time.Second)
// }

// // sends data to the channel
// func sell(ch chan string) {
// 	ch <- "Furniture"
// 	fmt.Println("Sent data to the channel")

// }

// // receive data from the channel
// func buy(ch chan string) {
// 	fmt.Println("Waiting for data")
// 	val := <-ch
// 	fmt.Println("Received data - ", val)

// }

//Buffered channels and how the send operation is blocked after the buffer gets full

//  func main() {
// 	var wg sync.WaitGroup
// 	// adding a counter of two since we are going to create two goroutines
// 	wg.Add(2)
// 	ch := make(chan int, 3)
// 	go sell(ch, &wg)
// 	// let the wait group wait and block the main function until all the goroutines return
// 	wg.Wait()
//  }

//  func sell(ch chan int, wg *sync.WaitGroup){
// 	ch <- 11
// 	ch <- 12
// 	ch <- 13
// 	go buy(ch, wg)// receives values from the channel
// 	ch <- 14 //extra line

// 	fmt.Println("Sent all data to the channel")
// 	wg.Done()

//  }

//  func buy(ch chan int, wg *sync.WaitGroup){
// 	fmt.Println("Waiting for data")
// 	//printing the data we recived
//     fmt.Println("Received data: ", <-ch)
// 	wg.Done()
//  }

//  // in the above, the buffer channel did not block the sell go-routine and this is because we did not exceed the limit
//  //if you send one more value on the channel we will get an error
//  // if we call the buy go-routine before the line that creates the deadlock, it wont give any error
//  //because the buy routine got called before the extra line was added
//  // if you also dont send anything from the sell channel there will also be a deadlock(because the sell channel is empty)

//Closing a channel
// func main(){
// 	ch := make(chan int, 10)
// 	ch <- 10
// 	ch <- 11
// 	val, ok := <-ch//receive first value from channel
// 	fmt.Println(val, ok)
// 	close(ch)
// 	val, ok = <-ch // receive second value from channel(:= was not used here because we already created the variable above)
// 	fmt.Println(val, ok)
// 	val, ok = <-ch
// 	fmt.Println(val, ok)// this returns 0 and false because 0 is the sero integer of int and false because there are no more values to receive in the go routine
// }

//PANIC SITUATION
//1. Sendind to a channel after it has been closed
//2. Closing an already closed channel

// func main(){
// 	ch := make(chan int, 10)
// 	ch <- 10
// 	ch <- 11
// 	val, ok := <-ch
// 	fmt.Println(val, ok)
// 	close(ch)
// 	ch <- 12
// }

//CHANNELS FOR-RANGE
// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	ch := make(chan int, 3)
// 	go sell(ch, &wg)
// 	go buy(ch, &wg)

// 	wg.Wait()

// }

// func sell(ch chan int, wg *sync.WaitGroup){
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println("Sent all the data")
// 	close(ch)
// 	wg.Done()
// }

// func buy(ch chan int, wg *sync.WaitGroup){
// 	fmt.Println("Waiting for the data")
// 	for val := range ch {
// 		fmt.Println("Received: ", val)
// 	}

// 	wg.Done()
// }

//USING CHANNEL FOR-RANGE IN BUFFERED CHANNELS
// func main(){
// 	ch := make(chan int, 5)

// 	// sending to the channel
// 	ch <- 100
// 	ch <- 200
// 	close(ch)

// 	// receiving from the channel using for-range
// 	for val := range ch{
// 		fmt.Println(val)
// 	}
// }
// above we can see that it is possible to close a channel but still
//have the remaining values be received from the same main function or in some other go routine
// in the last closing example, the value from received first before the channel was closed