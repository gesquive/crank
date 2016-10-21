package cmd

import "github.com/spf13/cobra"
import "github.com/spf13/viper"
import "github.com/gesquive/passmakr/passgen"

// xkcdCmd represents the xkcd command
var xkcdCmd = &cobra.Command{
	Use:     "xkcd",
	Aliases: []string{"x", "xk"},
	Short:   "Generate passwords with the XKCD password scheme",
	Long:    `Generate passwords with the XKCD password scheme`,
	Run:     runXKCDCmd,
}

func init() {
	RootCmd.AddCommand(xkcdCmd)

	xkcdCmd.Flags().IntP("words", "w", 4, "The number of words in the password")

	viper.BindPFlag("xkcd.words", xkcdCmd.Flags().Lookup("words"))

	viper.SetDefault("xkcd.words", 4)
}

func runXKCDCmd(cmd *cobra.Command, args []string) {
	numWords := viper.GetInt("xkcd.words")
	numPasses := viper.GetInt("number")
	passwords := passgen.GenerateXKCDCommonPasswords(numWords, numPasses)
	printList(passwords)
}
