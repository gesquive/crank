package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var displayVersion string

var showVersion bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:    "crank",
	Short:  "Generate random passwords",
	Long:   `Generate secure random passwords in a number of different formats`,
	Run:    runRootCmd,
	Hidden: true,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	displayVersion = version
	RootCmd.SetHelpTemplate(fmt.Sprintf("%s\nVersion:\n  github.com/gesquive/%s\n",
		RootCmd.HelpTemplate(), displayVersion))
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		"config file (default is $HOME/.config/crank.yml)")
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false,
		"Show the version and exit")
	RootCmd.PersistentFlags().Int32P("number", "n", 5, "The number of passwords to generate")

	viper.SetEnvPrefix("crank")
	viper.AutomaticEnv()

	viper.BindEnv("number")

	viper.BindPFlag("number", RootCmd.PersistentFlags().Lookup("number"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("config")              // name of config file (without extension)
	viper.AddConfigPath(".")                   // addign current director as first search path
	viper.AddConfigPath("$HOME/.config/crank") // adding home directory as next search path
	viper.AutomaticEnv()                       // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// crank random --chars="ABCD" --no-lower --no-upper --no-nums --no-special
	// crank xkcd --words=4
	// crank leet
	// crank trigraph
}

func runRootCmd(cmd *cobra.Command, args []string) {
	if showVersion {
		// cli.Info(displayVersion)
		fmt.Println(displayVersion)
		os.Exit(0)
	}
	cmd.Usage()
}

func printList(list []string) {
	for i := range list {
		fmt.Printf("%s\n", list[i])
	}
}
