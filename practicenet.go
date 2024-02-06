package main

import (
	"fmt"
	"strconv"
	"strings"
)


func scanMultiple(rports string) {
	var portslice []int

	portint := strings.Split(rports, "-")

	// concert port string to integer 
	startport, err:= strconv.Atoi(portint[0])

	if err == nil {
		fmt.Println(startport)
	}
	endport, err:= strconv.Atoi(portint[1])

	if err == nil {
		fmt.Println(endport)
	}

	//loop through and load it into a slice 
	for i := startport; i <= endport; i++ {
		portslice = append(portslice, i)
	}

	fmt.Println(portslice)

}

func main(){

	newstring := "1-10"
	scanMultiple(newstring)
}




// var portslice []int
// // concert port string to integer 
// portint := strings.Split(rports, "-")

// //convert the individual parts to integers
// startport:= strconv.Atoi(portint[0])
// endport:= strconv.Atoi(portint[1])

// //loop through and load it into a slice 
// for i := startport, i <= endport, i++ {
// 	portslice = append(portslice, i)
// }
// return portslice