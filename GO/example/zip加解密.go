package main

import (
	"bytes"
	"github.com/alexmullins/zip"
	"io"
	"log"
	"os"
)

func main() {
	contents := []byte("Hello World")

	// write a password zip
	raw := new(bytes.Buffer)
	zipw := zip.NewWriter(raw)
	w, err := zipw.Encrypt("hello.txt", "golang")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Close()

	// read the password zip
	zipr, err := zip.NewReader(bytes.NewReader(raw.Bytes()), int64(raw.Len()))
	if err != nil {
		log.Fatal(err)
	}
	for _, z := range zipr.File {
		z.SetPassword("golang")
		rr, err := z.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, rr)
		if err != nil {
			log.Fatal(err)
		}
		rr.Close()
	}
}
