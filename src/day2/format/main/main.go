package main

import "fmt"
type Human struct {
	Name string
	Age int
}

var people = Human{Name:"zhangsan",Age:90}
func main() {
	//testBasic()
	fmt.Printf("%v",people)
	fmt.Printf("%+v",people)
	fmt.Printf("%#v",people)
	fmt.Printf("%T",people)
	fmt.Printf("%%",people)

}

func testBasic() {
	var a int = 100
	var b bool
	c := 'a'
	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("90%%\n")
	fmt.Printf("%t\n", b)
	fmt.Printf("%b\n", 100)
	fmt.Printf("%f\n", 199.22)
	fmt.Printf("%q\n", "this is a test")
	fmt.Printf("%x\n", 39839333)
	fmt.Printf("%p\n", &a)
	str := fmt.Sprintf("a=%d", a)
	fmt.Printf("%q\n", str)
}
