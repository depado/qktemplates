package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"{{ .gitserver }}/{{ .organization }}/{{ .name }}/cmd"
	"{{ .gitserver }}/{{ .organization }}/{{ .name }}/infra"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// Main command that will be run when no other command is provided on the
// command-line
var rootCmd = &cobra.Command{
	Use:   "{{ .name }}",
	{{ if .description }}Short: "{{ .description }}",{{ end }}
	Run:   func(cmd *cobra.Command, args []string) { run() }, // nolint: unparam
}

// Version command that will display the build number and version (if any)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Printf("Build: %s\nVersion: %s\n", Build, Version) }, // nolint: unparam
}

func run() {
	// Setup CORS
	corsc := infra.NewCorsConfig(
		viper.GetBool("server.cors.enable"),
		viper.GetBool("server.cors.all"),
		viper.GetStringSlice("server.cors.origins"),
		viper.GetStringSlice("server.cors.methods"),
		viper.GetStringSlice("server.cors.headers"),
		viper.GetStringSlice("server.cors.expose"),
	)
	s := infra.NewServer(
		viper.GetString("server.host"),
		viper.GetInt("server.port"),
		viper.GetString("server.mode"),
		corsc, 		
		{{- if .prometheus }}
		viper.GetBool("prometheus.disabled"),
		viper.GetString("prometheus.prefix"),
		{{- end }}
	)
	s.Router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"alive": true}) })
	s.Start()
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
