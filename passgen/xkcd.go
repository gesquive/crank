package passgen

import "strings"

// GenerateXKCDPasswords generates a number of XKCD passwords with defined number of words
func GenerateXKCDPasswords(words int, number int, dictionary []string) []string {
	l := make([]string, number)
	for i := range l {
		l[i] = GenerateXKCDPassword(words, dictionary)
	}
	return l
}

// GenerateXKCDPassword generates passwords
func GenerateXKCDPassword(words int, dictionary []string) string {
	rands := getRandomUInt32Array(words)
	store := make([]string, words)
	for i := range store {
		store[i] = dictionary[rands[i]%uint32(len(dictionary))]
	}

	return strings.Join(store, " ")
}
