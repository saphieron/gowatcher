package cli

import (
	"strings"
	"time"

	"github.com/spf13/pflag"
)

type CliFlags struct {
	Interval time.Duration
	Version  bool
	Help     bool
	Verbose  bool
}

func ParseFlags() (*CliFlags, []string) {
	flagInterval := pflag.Float32P("interval", "n", 2.0, "seconds to wait between updates, in steps of 0.1")
	flagVersion := pflag.BoolP("version", "v", false, "display version information and exits.")
	flagHelp := pflag.BoolP("help", "h", false, "display this help and exits")
	flagVerbose := pflag.BoolP("verbose", "V", false, "report additional debugging information to STDERR")

	pflag.Parse()

	flags := &CliFlags{
		Interval: time.Duration(*flagInterval * float32(time.Second)),
		Version:  *flagVersion,
		Help:     *flagHelp,
		Verbose:  *flagVerbose,
	}
	commandPart := pflag.Args()
	if len(commandPart) == 1 {
		singleElementCommand := commandPart[0]
		if strings.HasPrefix(singleElementCommand, "\"") && strings.HasSuffix(singleElementCommand, "\"") {
			singleElementCommand = singleElementCommand[1:]
			singleElementCommand = singleElementCommand[:len(singleElementCommand)-1]
		}
		commandPart = tokenize(singleElementCommand)
	}
	return flags, commandPart
}

func PrintHelp() {
	pflag.Usage()
}

func tokenize(s string) []string {
	quoted := false

	return strings.FieldsFunc(s, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return !quoted && r == ' '
	})

}
