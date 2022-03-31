package wordlists

import (
	"bufio"
	"embed"
	"io"
	"os"
)

//go:embed assets/american-common
//go:embed assets/american-full
//go:embed assets/british-full
var dicts embed.FS

// LoadDictionaryFile loads a line delimited word list from a file
func LoadDictionaryFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readDictionary(file)
}

// LoadDictionaryAsset loads an internal word list
func LoadDictionaryAsset(assetName string) ([]string, error) {
	file, err := dicts.Open("assets/" + assetName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readDictionary(file)
}

func readDictionary(r io.Reader) (wordList []string, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return wordList, nil
}
