package passgen

import "io"
import "os"
import "bufio"
import "bytes"

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
	data, err := Asset(assetName)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)
	return readDictionary(reader)
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
