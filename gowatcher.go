package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/saphieron/gowatcher/cli"
	"github.com/saphieron/gowatcher/logging"
	"github.com/saphieron/gowatcher/loop"
)

func main() {
	Run()
	// Render()
}

func Run() {

	flags, rest := cli.ParseFlags()
	logging.Init(flags.Verbose)
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

	ctx, _ := prepareSignalCancelContex()
	looper := loop.NewLooper(flags.Interval, ctx)
	err := looper.Do(command, commandArgs...)
	if err != nil {
		slog.Error("Looper ran into error", "error", err.Error())
		os.Exit(1)
	}
	slog.Info("Exiting")
}

func printVersion() {
	fmt.Println("TODO: make a help message")
}

func prepareSignalCancelContex() (context.Context, context.CancelFunc) {
	servChannel := make(chan os.Signal, 1)
	signal.Notify(servChannel, os.Interrupt)
	newCtx, cancel := context.WithCancel(context.Background())
	go func() {
		<-servChannel
		slog.Warn("Received os interrupt")
		cancel()
	}()
	return newCtx, cancel
}
