package loop

import (
	"fmt"

	"github.com/saphieron/gowatcher/cli"
)

func Run() {
	fmt.Println("hello gowatcher")

	flags, rest := cli.ParseFlags()
	if flags.Version {
		printVersion()
		return
	}

	fmt.Printf("%v\n", flags)
	fmt.Printf("rest args %v\n", rest)
}

func printVersion() {
	fmt.Println("gowatcher")
	fmt.Printf("Version: '%s'\n", Version)
	fmt.Printf("BuildInfo: \n%s\n", BuildInfo)
}
