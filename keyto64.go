package main

import (
	"log"
	"os"
)

import (
	"github.com/eyedeekay/keyto/convert"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Please provide a full destination key as an argument")
	}
	b32, err := i2pdest.Base32(os.Args[1])
	if err != nil {
		panic(err)
	}
	b64, err := i2pdest.Base64(os.Args[1])
	if err != nil {
		panic(err)
	}
	log.Println("base32:", b32, "\nbase64:", b64)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
