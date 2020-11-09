package main

import (
	"io"
	"net/http"
)
func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"<h1> This is my firstPage </h1>")

}
func main() {
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8000",nil)
	
}
