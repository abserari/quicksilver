package main

import (
	"fmt"

	"gonum.org/v1/hdf5"
)

func main() {
	fmt.Println("=== go-hdf5 ===")
	version, err := hdf5.LibVersion()
	if err != nil {
		fmt.Printf("** error ** %s\n", err)
		return
	}
	fmt.Printf("=== version: %s", version)
	fmt.Println("=== bye.")
}
