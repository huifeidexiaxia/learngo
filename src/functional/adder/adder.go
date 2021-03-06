package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	// a := adder() is trivial and also works.
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		//s, a = a(i)
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d \n",
			i, s)
		pointer := reflect.ValueOf(a).Pointer()
		name := runtime.FuncForPC(pointer).Name()
		fmt.Println("the func name is %s, address is %x", name, a)
	}
}
