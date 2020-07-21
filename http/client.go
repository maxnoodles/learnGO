package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 8.1.0; V1818A Build/OPM1.171019.026; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/62.0.3202.84 Mobile Safari/537.36 SP-engine/2.16.0 baiduboxapp/11.20.0.14 (Baidu; P1 8.1.0)")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", s)
	fmt.Println(string(s))
}
