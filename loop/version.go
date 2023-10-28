package loop

import (
	"runtime/debug"
	"strings"
)

var Version = "development"

var BuildInfo = func() string {
	var buildInfo strings.Builder
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			buildInfo.WriteString("\t")
			buildInfo.WriteString(setting.Key)
			buildInfo.WriteString("-")
			buildInfo.WriteString(setting.Value)
			buildInfo.WriteString("\n")
		}
	}
	return buildInfo.String()
}()

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}()
