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

// GenerateTrigraphDictionary generates a trigraph dictionary from a word list
func GenerateTrigraphDictionary(dictionary []string) (sigma int, tris [26][26][26]int) {
	for w := 0; w < len(dictionary); w++ {
		k1, k2 := -1, -1
		for c := 0; c < len(dictionary[w]); c++ {
			k3 := getRuneIndex(dictionary[w][c])
			if k3 >= 0 { // we have a valid character
				if k1 >= 0 { // do we have 3 letters?
					tris[k1][k2][k3]++ // count
					sigma++
				}
				k1 = k2 // shift over
				k2 = k3
			}
		}
	}
	return sigma, tris
}

func getRuneIndex(char byte) int {
	if char >= 'A' && char <= 'Z' {
		return int(char - 'A')
	} else if char >= 'a' && char <= 'z' {
		return int(char - 'a')
	}
	return -1
}
