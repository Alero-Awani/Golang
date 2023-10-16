package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
		//each of the case statements is waiting for the receive operation to complete 
		//on its respective channel// the output is non-deterministic //even if the go routines execute at the same time, the select statement might just choose one of them randomly
		case val1 := <- ch1: //val1 is receiving values from channelone
			fmt.Println(val1)
		case val2 := <- ch2:
			fmt.Println(val2)
			// if no goroutine sends to the channel, the default case gets executed 
		default:
			fmt.Println("Executed default block")
	}
	time.Sleep(1 * time.Second)

}


func goOne(ch1 chan string){
	//ch1 <- "Channel-1"

}

func goTwo(ch2 chan string){
	//ch2 <- "Channel-2"
}