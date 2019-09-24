package main

import (
	transloadit "gopkg.in/transloadit/go-sdk.v1"
)

func main() {
	// Create client
	options := transloadit.DefaultConfig
	options.AuthKey = "b211384720824c36ae7e844f4d2cb0c1"
	options.AuthSecret = "2c43615948c97542356e030685a451a95b3e09d0"

	client := transloadit.NewClient(options)

	assembly := transloadit.NewAssembly()

}
