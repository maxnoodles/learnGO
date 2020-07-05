package main

import (
	"fmt"
	"learnGo/retriever/mock"
	"learnGo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

type Sessions interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func session(s Sessions) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func inspect(r Retriever) {
	fmt.Println(r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(">>contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println(">>UserAgent: ", v.UserAgent)
	}
}

func main() {
	r := &mock.Retriever{Contents: "fake contents"}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
	r2 := &real.Retriever{
		UserAgent: "chrome",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r2, r2)

	inspect(r)
	inspect(r2)

	// type assertion
	if realRetriever, ok := interface{}(r2).(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a real retriever")
	}

	fmt.Println(session(r))
	//fmt.Println(download(r2))
}
