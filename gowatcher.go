package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/saphieron/gowatcher/cli"
	"github.com/saphieron/gowatcher/logging"
	"github.com/saphieron/gowatcher/loop"
)

func main() {
	Run()
}

func Run() {
	logging.Init()

	flags, rest := cli.ParseFlags()
	if flags.Help {
		cli.PrintHelp()
		return
	}
	if flags.Version {
		printVersion()
		return
	}
	command := rest[0]
	var commandArgs []string
	if len(rest) > 1 {
		commandArgs = rest[1:]
	}

	ctx := prepareSignalCancelContex()
	looper := loop.NewLooper(flags.Interval, ctx)
	err := looper.Start(command, commandArgs)
	if err != nil {
		os.Exit(1)
	}
}
func printVersion() {
	fmt.Println("TODO: make a help message")
}

func prepareSignalCancelContex() context.Context {
	servChannel := make(chan os.Signal, 1)
	signal.Notify(servChannel, os.Interrupt)
	newCtx, cancel := context.WithCancel(context.Background())
	go func() {
		<-servChannel
		logging.ErrorLog.Print("cancelling via interrupt")
		cancel()
	}()
	return newCtx
}
