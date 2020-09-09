package main

import (
	"fmt"

	"github.com/abserari/quicksilver/gosample/export-unexportstructure/unexport"
)

func main() {
	v := unexport.ExportV()
	p := unexport.ExportP()

	// p.name or v.name is not true because
	// name are not export
	fmt.Println(v.Name(), p.Name(), " ", v.Name(), " ", p.name)
}
