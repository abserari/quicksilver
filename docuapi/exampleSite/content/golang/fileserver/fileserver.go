package main

import "net/http"

// net/http
func main() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
}

// func main() {
// 	http.ListenAndServe(":8081", http.FileServer(http.Dir("./")))
// }

//gin
// func main() {
// 	router := gin.Default()
// 	router.StaticFS("/", gin.Dir("./gin/", true))
// 	router.Run(":8080")
// }
