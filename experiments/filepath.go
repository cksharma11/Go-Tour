package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	filename := filepath.Base("./")
	fmt.Println(filename)
}