package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gesquive/crank/passgen"
	"github.com/gesquive/crank/wordlists"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

	xkcdCmd.Flags().IntP("words", "w", 4, "The number of words in the password.")
	xkcdCmd.Flags().StringP("dictionary", "d", "", "Word dictionary to generate passwords from.")
	xkcdCmd.Flags().StringP("separator", "s", " ", "Separator to use between words.")

	viper.BindPFlag("xkcd.words", xkcdCmd.Flags().Lookup("words"))
	viper.BindPFlag("xkcd.dictionary", xkcdCmd.Flags().Lookup("dictionary"))
	viper.BindPFlag("xkcd.separator", xkcdCmd.Flags().Lookup("separator"))

	viper.SetDefault("xkcd.words", 4)
	viper.SetDefault("xkcd.dictionary", "asset:american-common")
}

func runXKCDCmd(cmd *cobra.Command, args []string) {
	numWords := viper.GetInt("xkcd.words")
	numPasses := viper.GetInt("number")
	dictName := viper.GetString("xkcd.dictionary")
	separator := viper.GetString("xkcd.separator")

	var dictionary []string
	var err error
	if strings.HasPrefix(dictName, "asset:") {
		dictName = strings.TrimPrefix(dictName, "asset:")
		dictionary, err = wordlists.LoadDictionaryAsset(dictName)
	} else {
		dictionary, err = wordlists.LoadDictionaryFile(dictName)
	}
	if err != nil {
		fmt.Printf("loading error: %+v\n", err)
		os.Exit(1)
	}

	passwords := passgen.GenerateXKCDPasswords(numWords, numPasses, separator, dictionary)
	printList(passwords)
}
