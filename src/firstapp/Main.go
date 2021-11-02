package main

// Woran manchmal das laden scheitert: https://stackoverflow.com/questions/37328834/how-to-specify-the-file-location-for-template-parsefiles-in-go-language !!
// Conditional Rendering https://stackoverflow.com/questions/29689426/conditional-rendering-of-html-in-golang-layout-tpl-by-session-variable
import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

type User struct {
	Username string
	Password string
	// 0 = Schreiberlaubnis || 1 = Admin
	AccessLevel int
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	timon := User{Username: "timonsrm", Password: "meinpasswort", AccessLevel: 1}
	t := template.Must(template.ParseFiles("login.html.tmpl"))
	t.Execute(w, timon)
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
	http.HandleFunc("/login/", loginHandler)
	http.ListenAndServe(":80", nil)
}
