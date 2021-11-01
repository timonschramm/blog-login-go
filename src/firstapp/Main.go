package main

//https://stackoverflow.com/questions/37328834/how-to-specify-the-file-location-for-template-parsefiles-in-go-language !!
import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	fmt.Println(p)
	t := template.Must(template.ParseFiles("basictemplating.html"))
	fmt.Println(t)
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello World </h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":80", nil)
}
