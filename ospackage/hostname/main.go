package main

import (
	"fmt"
	"os"
)

func main() {
host, _ := os.Hostname()
fmt.Println(host)
}