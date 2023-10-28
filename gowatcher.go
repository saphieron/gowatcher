package main

import (
	"fmt"

	"github.com/saphieron/gowatcher/loop"
)

var Version = "maindev"

func main() {
	fmt.Println(Version)
	loop.Run()
}
