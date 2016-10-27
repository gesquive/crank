package cmd

import "os"
import "fmt"
import "github.com/gesquive/passforge/passgen"
import "github.com/spf13/cobra"
import "github.com/spf13/viper"

// trigraphCmd represents the trigraph command
var trigraphCmd = &cobra.Command{
	Use:     "trigraph",
	Aliases: []string{"t", "tri"},
	Short:   "Generate human readable trigraph passwords",
	Long:    `Generate human readable trigraph based passwords from a word list`,
	Run:     runTrigraphCmd,
}

func init() {
	RootCmd.AddCommand(trigraphCmd)

	trigraphCmd.Flags().IntP("length", "l", 16, "Length of password")
	trigraphCmd.Flags().StringP("dictionary", "d", "", "Word dictionary to generate trigraphs from.")

	viper.BindPFlag("trigraph.length", trigraphCmd.Flags().Lookup("length"))
	viper.BindPFlag("trigraph.dictionary", trigraphCmd.Flags().Lookup("dictionary"))

	viper.SetDefault("trigraph.length", 16)

}

func runTrigraphCmd(cmd *cobra.Command, args []string) {
	length := viper.GetInt("trigraph.length")
	numPasses := viper.GetInt("number")
	dictName := viper.GetString("trigraph.dictionary")
	var dictionary []string
	var err error
	if len(dictName) == 0 {
		dictionary, err = passgen.LoadDictionaryAsset("american-full")
	} else {
		dictionary, err = passgen.LoadDictionaryFile(dictName)
	}
	if err != nil {
		fmt.Printf("ERR: %+v\n", err)
		os.Exit(1)
	}

	sigma, tris := passgen.GenerateTrigraphDictionary(dictionary)

	_ = "breakpoint"
	passwords := passgen.GenerateTrigraphPasswords(length, numPasses, sigma, tris)
	printList(passwords)

}
