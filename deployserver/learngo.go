package main

import (
	"io"
	"net/http"
)

func Print1To20() int {
	res := 0;
	for i := 1; i<=20 ; i++{
		res += i
	}
	return res
}

func firstPage(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1> This is my first page</h1>");

}

func main() {
	http.HandleFunc("/",firstPage);
	http.ListenAndServe(":8001",nil);
}
