---
if: cli
---
package cmd

import (
	"strings"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AddLoggerFlags adds support to configure the level of the logger
func AddLoggerFlags(c *cobra.Command) {
	c.PersistentFlags().String("log.level", "info", "one of debug, info, warn, error or fatal")
	c.PersistentFlags().String("log.format", "text", "one of text or json")
	c.PersistentFlags().Bool("log.line", false, "enable filename and line in logs")
}

// AddConfigurationFlag adds support to provide a configuration file on the
// command line
func AddConfigurationFlag(c *cobra.Command) {
	c.PersistentFlags().String("conf", "", "configuration file to use")
}

// AddAllFlags will add all the flags provided in this package to the provided
// command and will bind those flags with viper
func AddAllFlags(c *cobra.Command) {
	AddConfigurationFlag(c)
	AddLoggerFlags(c)
	if err := viper.BindPFlags(c.PersistentFlags()); err != nil {
		logrus.WithError(err).WithField("step", "AddAllFlags").Fatal("Couldn't bind flags")
	}
}

// Initialize will be run when cobra finishes its initialization
func Initialize() {
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("{{ .name }}")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Configuration file
	if viper.GetString("conf") != "" {
		viper.SetConfigFile(viper.GetString("conf"))
	} else {
		viper.SetConfigName("conf")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/config/")
	}
	hasconf := viper.ReadInConfig() == nil

	// Set log level
	lvl := viper.GetString("log.level")
	if l, err := logrus.ParseLevel(lvl); err != nil {
		logrus.WithFields(logrus.Fields{"level": lvl, "fallback": "info"}).Warn("Invalid log level")
	} else {
		logrus.SetLevel(l)
	}

	// Set log format
	switch viper.GetString("log.format") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	// Defines if logrus should display filenames and line where the log ocured
	if viper.GetBool("log.line") {
		logrus.AddHook(filename.NewHook())
	}

	// Delays the log for once the logger has been setup
	if hasconf {
		logrus.WithField("file", viper.ConfigFileUsed()).Debug("Found configuration file")
	} else {
		logrus.Debug("No configuration file found")
	}
}
