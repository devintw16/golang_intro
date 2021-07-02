package main

import (
	"fmt"
	"net/http"
)

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Thank you for contacting us, we will be in touch with you soon!</h1>")
	fmt.Println("The following person is reaching out to us:")
    r.ParseForm()
    fmt.Println("",r.Form["name"])
    fmt.Println("",r.Form["email"])
    fmt.Println("",r.Form["message"])
}

func main() {
        http.Handle("/", http.FileServer(http.Dir("../html")))
        http.HandleFunc("/contact", contact )
        http.ListenAndServe(":8000", nil)
	
}