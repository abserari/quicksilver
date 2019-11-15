package main
 
import "fmt"
 
func decorator(f func(s string)) func(s string) {
 
        return func(s string) {
                fmt.Println("Started")
                f(s)
                fmt.Println("Done")
        }
}
 
func Hello(s string) {
        fmt.Println(s)
}
 
func main() {
        decorator(Hello)("Hello, World!")
}