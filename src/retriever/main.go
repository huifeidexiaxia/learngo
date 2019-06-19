package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

type Parse interface {
	Parse()
} 
const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	//var r Retriever
	retriever := mock.Retriever{}
	retriever.Parse()

	retriever2 := real.Retriever{}
	retriever2.Parse()
	
	fmt.Print(retriever)
	fmt.Print(retriever2)

	//mockRetriever := mock.Retriever{
	//	Contents: "this is a fake imooc.com"}
	//r = &mockRetriever
	//inspect(r)
	//
	//r = &real.Retriever{
	//	UserAgent: "Mozilla/5.0",
	//	TimeOut:   time.Minute,
	//}
	//
	//retriever := mockRetriever
	//retriever.Get("")
	//inspect(r)
	//
	//// Type assertion
	//if mockRetriever, ok := r.(*mock.Retriever); ok {
	//	fmt.Println(mockRetriever.Contents)
	//} else {
	//	fmt.Println("r is not a mock retriever")
	//}
	//
	//fmt.Println(
	//	"Try a session with mockRetriever")
	//fmt.Println(session(&mockRetriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
