// Copyright © 2019 Metrum Research Group
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/metrumresearchgroup/cct/configlib"
	"github.com/spf13/cobra"
)

// VERSION is the current pkc version
const VERSION string = "0.0.1-alpha.1"

var cfg configlib.Config

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cct",
	Short: "conventional commits",
	Long:  fmt.Sprintf("cct cli version %s", VERSION),
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().String("config", "", "config file (default is cct.yml)")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))

	RootCmd.PersistentFlags().String("loglevel", "", "level for logging")
	viper.BindPFlag("loglevel", RootCmd.PersistentFlags().Lookup("loglevel"))

	RootCmd.PersistentFlags().Bool("preview", false, "preview action, but don't actually run command")
	viper.BindPFlag("preview", RootCmd.PersistentFlags().Lookup("preview"))

	RootCmd.PersistentFlags().Bool("debug", false, "use debug mode")
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
}

func setGlobals() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	switch logLevel := strings.ToLower(viper.GetString("loglevel")); logLevel {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" { // enable ability to specify config file via flag
	// 	viper.SetConfigFile(cfgFile)
	// }
	configlib.LoadConfigFromPath(viper.GetString("config"))

	setGlobals()
	if viper.GetBool("debug") {
		viper.Debug()
	}
	viper.Unmarshal(&cfg)
	configDir, _ := filepath.Abs(viper.ConfigFileUsed())
	cwd, _ := os.Getwd()
	log.WithFields(logrus.Fields{
		"cwd": cwd,
		"nwd": filepath.Dir(configDir),
	}).Trace("setting directory to configuration file")
	// TODO: don't change path if the config was read in from the users home directory
	// as they might use a global configuration set but would want to stay in the repo
	os.Chdir(filepath.Dir(configDir))
}
