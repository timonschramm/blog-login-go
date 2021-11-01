package main

import (
	"fmt"
	"net/http"
)

func index_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<p> this is HTML</p>")
	fmt.Fprintf(w, "You can %s add %s", "even", "something cool")

	fmt.Fprintf(w, `<h1>Hey there</h1>
					<p> this is HTML</p>
	`)
}
func main() {
	http.HandleFunc("/", index_Handler)
	http.ListenAndServe(":8000", nil)
}

func main2() {
	fmt.Println("Enter your name:")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Dein Name ist also " + first)

	fmt.Println("Bitte gib dein Username ein:")
	var zUsername string
	fmt.Scanln(&zUsername)
	fmt.Println("Bitte gib dein Passwort ein.")
	var zPassword string
	fmt.Scanln(&zPassword)
	fmt.Println("Hallo " + zUsername + " dein Passwort ist: " + zPassword)
}
