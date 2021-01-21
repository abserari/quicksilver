package main

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "aF_QnX9Zz4iRVH0ed5YgDtGTOCrlGl6Bs19lRT5vfgmsILsm5YJJD7ZbPZpdcLXYUWOzoWo4Iy2zm2fgLjS27w=="
	const bucket = "bucket"
	const org = "abser"

	client := influxdb2.NewClient("http://192.168.0.254:8086", token)
	// always close client at the end
	defer client.Close()

	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 22.5, 45.0))
	// Flush writes
	writeAPI.Flush()

	query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"stat\")", bucket)
	// Get query client
	queryAPI := client.QueryAPI(org)
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}
