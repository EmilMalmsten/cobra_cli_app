package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var GoBlueprintVersion string

func getGoBlueprintVersion() string {
	if len(GoBlueprintVersion) != 0 {
		return GoBlueprintVersion
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "unknown"
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information.",
	Long: `The version command provides information about the application's version.
Use this command to check the current version of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		goBlueprintVersion := getGoBlueprintVersion()
		fmt.Printf("cobra cli version %v\n", goBlueprintVersion)
	},
}
