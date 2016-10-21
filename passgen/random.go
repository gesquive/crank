package passgen

// GenerateDefaultRandomPasswords generates a number of configurable random passwords
func GenerateDefaultRandomPasswords(length int, number int, useUpper bool,
	useLower bool, useNumeric bool, useSpecial bool) []string {

	var lowerAlpha = []byte("abcdefghijklmnopqrstuvwxyz")
	var upperAlpha = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numeric = []byte("0123456789")
	var special = []byte("!$%@#")

	runes := []byte{}
	if useLower {
		runes = append(runes, lowerAlpha...)
	}
	if useUpper {
		runes = append(runes, upperAlpha...)
	}
	if useNumeric {
		runes = append(runes, numeric...)
	}
	if useSpecial {
		runes = append(runes, special...)
	}

	l := make([]string, number)
	for i := range l {
		l[i] = GenerateRandomPassword(length, runes)
	}
	return l

}

// GenerateRandomPasswords generates a number of random passwords
func GenerateRandomPasswords(length int, number int, runes []byte) []string {
	l := make([]string, number)
	for i := range l {
		l[i] = GenerateRandomPassword(length, runes)
	}
	return l
}

// GenerateRandomPassword generates a random password
func GenerateRandomPassword(length int, runes []byte) string {
	b := make([]byte, length)
	rands := getRandomUInt16Array(length)
	for i := range b {
		b[i] = runes[rands[i]%uint16(len(runes))]
	}

	return string(b)
}
