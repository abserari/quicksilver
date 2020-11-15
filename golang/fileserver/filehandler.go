package main

import "net/http"

func main() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
}

// func main() {
// 	http.ListenAndServe(":8081", http.FileServer(http.Dir("./")))
// }
