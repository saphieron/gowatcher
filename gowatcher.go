package main

import (
	"fmt"

	"github.com/saphieron/gowatcher/cli"
)

func main() {
	fmt.Println("hello gowatcher")

	flags := cli.ParseFlags()
	fmt.Printf("%v\n", flags)
}
