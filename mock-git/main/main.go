package main

import (
	"fmt"
	"../targzhelpers"
)

func main() {
	targzhelpers.TarIt("../main", "../main")
	targzhelpers.UnTar("./.tars/main.tar", "./.tars/main")
	fmt.Println("done")
}