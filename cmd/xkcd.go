package cmd

import "os"
import "fmt"
import "github.com/spf13/cobra"
import "github.com/spf13/viper"
import "github.com/gesquive/passforge/passgen"

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
	xkcdCmd.Flags().StringP("dictionary", "d", "", "Word dictionary to generate passwords from.")

	viper.BindPFlag("xkcd.words", xkcdCmd.Flags().Lookup("words"))
	viper.BindPFlag("xkcd.dictionary", xkcdCmd.Flags().Lookup("dictionary"))

	viper.SetDefault("xkcd.words", 4)
}

func runXKCDCmd(cmd *cobra.Command, args []string) {
	numWords := viper.GetInt("xkcd.words")
	numPasses := viper.GetInt("number")
	dictName := viper.GetString("xkcd.dictionary")
	var dictionary []string
	var err error
	if len(dictName) == 0 {
		dictionary, err = passgen.LoadDictionaryAsset("american-common")
	} else {
		dictionary, err = passgen.LoadDictionaryFile(dictName)
	}
	if err != nil {
		fmt.Printf("ERR: %+v\n", err)
		os.Exit(1)
	}

	passwords := passgen.GenerateXKCDPasswords(numWords, numPasses, dictionary)
	printList(passwords)
}
