package main

import "net/http"

func main(){
	  http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(HomeFolder + "images/"))))
}
