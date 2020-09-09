package main

// a snippet from a student @Sun Mingyu
import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/upload", UploadFile)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// UploadFile -
func UploadFile(response http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	if request.Method == "GET" {
		// Generate token
		currentTime := time.Now().Unix()
		hash := md5.New()
		io.WriteString(hash, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		// ParseFiles Create a template object and parse files
		template, err := template.ParseFiles("upload.html")
		if err != nil {
			log.Fatal(err)
		}

		// Write token to response
		template.Execute(response, token)
	} else {
		// Control the size of the read memory
		request.ParseMultipartForm(32 << 20)
		// Gets the form file based on the field name
		formFile, handler, err := request.FormFile("uploadfile")
		if err != nil {
			log.Fatal(err)
		}
		defer formFile.Close()
		fmt.Fprintf(response, "%v", handler.Header)

		// Open save file
		file, err := os.OpenFile("test.go", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}

		// Read the form file and write to the save file
		_, err = io.Copy(file, formFile)
		if err != nil {
			log.Fatal(err)
		}
	}

}
