package cli

import (
	"github.com/spf13/pflag"
)

type CliFlags struct {
	Interval float32
	Version  bool
	Help     bool
}

func ParseFlags() *CliFlags {
	flagInterval := pflag.Float32P("interval", "n", 2.0, "seconds to wait between updates, in steps of 0.1")
	flagVersion := pflag.BoolP("version", "v", false, "display version information and exits.")
	flagHelp := pflag.BoolP("help", "h", false, "display this help and exits")
	pflag.Parse()

	flags := &CliFlags{
		Interval: *flagInterval,
		Version:  *flagVersion,
		Help:     *flagHelp,
	}
	return flags
}
