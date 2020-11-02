package main

import "C"

//export FunctionTest
func FunctionTest(pub string) *C.char {
	return C.CString("Hello " + pub)
}

func main() {
}
