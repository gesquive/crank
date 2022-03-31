package cmd

import (
	"github.com/gesquive/crank/passgen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:     "random",
	Aliases: []string{"r", "rdm"},
	Short:   "Generate random passwords",
	Long:    `Generate random passwords from character lists`,
	Run:     runRandomCmd,
}

func init() {
	RootCmd.AddCommand(randomCmd)

	randomCmd.Flags().IntP("length", "l", 24, "Length of password")
	randomCmd.Flags().BoolP("lower", "o", false, "Use lower case alpha characters")
	randomCmd.Flags().BoolP("upper", "u", false, "Use upper case alpha characters")
	randomCmd.Flags().BoolP("numeric", "m", false, "Use numeric characters")
	randomCmd.Flags().BoolP("special", "s", false, "Use special characters")
	randomCmd.Flags().StringP("chars", "C", "", "The characters to generate the password from, all other flags are ignored")

	viper.BindPFlag("random.length", randomCmd.Flags().Lookup("length"))
	viper.BindPFlag("random.lower", randomCmd.Flags().Lookup("lower"))
	viper.BindPFlag("random.upper", randomCmd.Flags().Lookup("upper"))
	viper.BindPFlag("random.numeric", randomCmd.Flags().Lookup("numeric"))
	viper.BindPFlag("random.special", randomCmd.Flags().Lookup("special"))
	viper.BindPFlag("random.chars", randomCmd.Flags().Lookup("chars"))

	viper.SetDefault("random.length", 24)
	viper.SetDefault("random.lower", false)
	viper.SetDefault("random.upper", false)
	viper.SetDefault("random.numeric", false)
	viper.SetDefault("random.special", false)
	viper.SetDefault("random.chars", "")
}

func runRandomCmd(cmd *cobra.Command, args []string) {
	length := viper.GetInt("random.length")
	numPasses := viper.GetInt("number")
	runes := []byte(viper.GetString("random.chars"))
	useLower := viper.GetBool("random.lower")
	useUpper := viper.GetBool("random.upper")
	useNumeric := viper.GetBool("random.numeric")
	useSpecial := viper.GetBool("random.special")
	var passwords []string
	if len(runes) > 0 {
		// Then ignore the rest of the flags
		passwords = passgen.GenerateRandomPasswords(length, numPasses, runes)
	} else if useLower || useUpper || useNumeric || useSpecial {
		// If any of these flags are set, use only the passed in flags
		passwords = passgen.GenerateDefaultRandomPasswords(length, numPasses, useUpper, useLower, useNumeric, useSpecial)
	} else {
		// By default use all the character sets
		passwords = passgen.GenerateDefaultRandomPasswords(length, numPasses, true, true, true, true)
	}
	printList(passwords)
}
