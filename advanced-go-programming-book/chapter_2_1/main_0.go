package main

//#include "hello.c"
import "C"

func main () {
	C.SayHello(C.CString("Hello, World\n"))
}