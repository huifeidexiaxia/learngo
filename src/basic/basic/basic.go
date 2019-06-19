package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
type Response struct {
	Code int64 `json:"code"`
	Msg string `json:"message"`
	Data interface{}
}
func main() {
	//var resp Response
	//err:=json.Unmarshal([]byte(` {"code":200,"message":"success","data":"Macbook Pro 2018 款"}`),&resp)
	//if err !=nil {
	//	return
	//}else {
	//	fmt.Println(resp)
	//}
	//
	//bytes, err := json.Marshal(resp)
	//s2 := string(bytes)
	//
	//fmt.Print(s2)
	//fmt.Println(s2==` {"code":200,"msg":"success","data":"Macbook Pro 2018 款"}`)
	//fmt.Println(s2==" {\"code\":200,\"msg\":\"success\",\"data\":\"Macbook Pro 2018 款\"}")
	//fmt.Println(`{"code":200,"msg":"success","data":"Macbook Pro 2018 款"}`)
	//fmt.Println("{\"code\":200,\"msg\":\"success\",\"data\":\"Macbook Pro 2018 款\"}")


	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
