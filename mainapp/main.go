package main

import (
	"fmt"
	"monorepo-release-notes/subapp"
)

func main() {
	fmt.Println(subapp.GetMessage())
	fmt.Println("Hello, ")
}