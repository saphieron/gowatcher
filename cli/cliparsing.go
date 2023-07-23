package cli

import (
	flag "github.com/spf13/pflag"
)

type CliFlags struct {
	interval float32
	version  bool
	help     bool
}

func ParseFlags() *CliFlags {
	flagInterval := flag.Float32P("interval", "n", 2.0, "seconds to wait between updates, in steps of 0.1")
	flagVersion := flag.BoolP("version", "v", false, "display version information and exits.")
	flagHelp := flag.BoolP("help", "h", false, "display this help and exits")
	// flag.Lookup("flagname").NoOptDefVal = "4321"
	flag.Parse()

	flags := &CliFlags{
		interval: *flagInterval,
		version:  *flagVersion,
		help:     *flagHelp,
	}
	return flags
}
