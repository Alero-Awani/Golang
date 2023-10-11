package main

import (
	"fmt"
	"sync"
	"time"
)


func calculateSquare(i int, wg *sync.WaitGroup){
	fmt.Println(i * i)
	// The done method will be called by each of the go routines, hence it will keep decreasing our counter by one for each go routine 
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	start := time.Now()
	//We want to wait for 10 go routines so we set the counter to ten
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go calculateSquare(i, &wg)
	}

	 elapsed := time.Since(start)
	 time.Sleep(2 * time.Second)
	 // We want to block the execution of the main go routine until all the other go routines execute hence wg.wait
	 wg.Wait()
	 fmt.Println("Function took ", elapsed)
}
//adding go to the calculatesquare function means that the 10,000 function calls will be executed in 10,000 different go routines 


// func main(){
// 	go start()
// 	time.Sleep(1 * time.Second)

// }

// func start(){
// 	go process()
// 	fmt.Println("In start..")
// }

// func process(){
// 	fmt.Println("in process..")
// }