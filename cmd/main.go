package main

import (
	"echo"
	"log"
)

func main() {
	err := echo.Run()
	if err != nil {
		log.Fatal(err)
	}
}
