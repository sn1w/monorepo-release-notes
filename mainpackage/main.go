package main

import (
	"fmt"
	"github.com/sn1w/monorepo-release-notes/subpackage"
)

func main() {
	fmt.Printf("Hello, %s\n", subpackage.GetSubpackageMessage())
}