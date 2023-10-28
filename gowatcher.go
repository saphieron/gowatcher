package main

import (
	"fmt"

	"github.com/saphieron/gowatcher/cli"
)

func main() {
	fmt.Println("hello gowatcher")

	flags, rest := cli.ParseFlags()

	fmt.Printf("%v\n", flags)
	fmt.Printf("rest args %v\n", rest)
}
