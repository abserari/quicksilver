package main

import "fmt"
import "regexp"

func main() {
	s := `<button style="hello" title="aaaa" id="bbb" class="ccc">`

	r, _ := regexp.Compile(`<button([^\<\>]+)>`)
	matches := r.FindStringSubmatch(s)

	fmt.Println(matches[0])
	fmt.Println(matches[1])
}
