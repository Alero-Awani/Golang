package main

import (
	"fmt"
)

type ByteCounter int

// the write method of the bytecounter type simply counts the bytes written to it before discarding them
func (c *ByteCounter) Write(p []byte) (int, error) {
	//convert int to ByteCounter , i.e  convert the length of the byte slice p into a ByteCounter.
	*c += ByteCounter(len(p))
	return len(p), nil
}

//the method above takes in a slice of bytes
//it increments the value of the bytecounter by the lenght of the slice

func main() {
	var c ByteCounter
	// c.Write([]byte("hello"))
	// fmt.Println(c)

	//reset the counter 
	var name = "Dolly"
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println(c)

}
//Since *ByteCounter satisfies the io.Writer contract, we can pass it to Fprintf,
//func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

//EXTRA 
//Source code for Fprintf looks like this
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
// 	p := newPrinter()
// 	p.doPrintf(format, a)
// 	n, err = w.Write(p.buf)
// 	p.free()
// 	return
//  }

 //So when you call Fprintf, internally Write is called.//to learn more about where newPrinter comes from - https://github.com/golang/go/blob/go1.16.2/src/fmt/print.go#L202-L208
 //what this basically means is that c.Write was called in the Fprintf function. So we can format the result of our byte counter using Fprintf because c has a write method and
 //Fprintf takes in an io.Writer interface

//  //COUNTING WRITER EXERCISE 


//  //Write a function CountingWriter with the signature below that, given an io.Writer, 
//  //returns a new Writer that wraps the original, and a pointer to an int64 variable that 
//  //at any moment contains the number of bytes written to the new Writer.

//  type CountingWriter struct {
// 	writer     		io.Writer
// 	bytesWritten	int
//  }

//  // New creates a new writer that wraps w.  The wrapping writer counts
// // the number of bytes written to the wrapped writer.

// //this is a constructor function that takes io.writer as a paramter and returns a new CountingWriter
// //then it initialises the Counting writer with w and 0
//  func New(w io.Writer) *CountingWriter {
// 	return &CountingWriter{
// 		writer: 		w,
// 		bytesWritten:	0,
// 	}
//  }



// func (w *CountingWriter) Write(b []byte) (int, error) {
// 	n, err := w.writer.Write(b)
// 	w.bytesWritten += n
// 	return n, err
// }

// //BytesWritten returns the number of bytes that was written to the wrapped writer
// func(w *CountingWriter) BytesWritten() int {
// 	return w.bytesWritten
// }


//COUNTING WRITER EXERCISE SIMPLIFIED  - https://xingdl2007.gitbooks.io/gopl-soljutions/content/chapter-7-interfaces.html

