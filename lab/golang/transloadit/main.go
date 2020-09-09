package main

import (
	"context"
	"fmt"

	transloadit "gopkg.in/transloadit/go-sdk.v1"
)

func main() {
	// Create client
	options := transloadit.DefaultConfig
	options.AuthKey = "b211384720824c36ae7e844f4d2cb0c1"
	options.AuthSecret = "2c43615948c97542356e030685a451a95b3e09d0"

	client := transloadit.NewClient(options)

	assembly := transloadit.NewAssembly()

	assembly.AddFile("image", "/Users/abser/Pictures/收集图/丛林狗武士.jpg")
	assembly.AddStep("resize", map[string]interface{}{
		"robot":           "/image/resize",
		"width":           75,
		"height":          75,
		"resize_strategy": "pad",
		"background":      "#000000",
	})

	info, err := client.StartAssembly(context.Background(), assembly)
	if err != nil {
		panic(err)
	}

	info, err = client.WaitForAssembly(context.Background(), info)
	if err != nil {
		panic(err)
	}

	fmt.Printf("You can view the result at %s\n", info.Results["resize"][0].SSLURL)
}
