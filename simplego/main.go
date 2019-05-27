package main
{{ if .cli }}
import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	
	"{{ .gitserver }}/{{ .organization }}/{{ .name }}/cmd"
)
{{ end }}
// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)
{{ if .cli }}
// Main command that will be run when no other command is provided on the
// command-line
var rootCmd = &cobra.Command{
	Use:   "{{ .name }}",
	Run:   func(cmd *cobra.Command, args []string) { run() }, // nolint: unparam
}

// Version command that will display the build number and version (if any)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Printf("Build: %s\nVersion: %s\n", Build, Version) }, // nolint: unparam
}

func run() {
	// Do your things
}

func main() {
	// Initialize Cobra and Viper
	cobra.OnInitialize(cmd.Initialize)
	cmd.AddAllFlags(rootCmd)
	rootCmd.AddCommand(versionCmd)

	// Run the command
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Couldn't start")
	}
}
{{ else }}
func main() {

}
{{ end }}
