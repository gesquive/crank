package passgen

// GenerateTrigraphPasswords creates a number of trigraph passwords
func GenerateTrigraphPasswords(length int, number int,
	sigma int, tris [26][26][26]int) []string {
	l := make([]string, number)
	for i := range l {
		l[i] = GenerateTrigraphPassword(length, sigma, tris)
	}
	return l
}

// GenerateTrigraphPassword creates a password from the defined sigma and trigraphs
func GenerateTrigraphPassword(length int, sigma int, tris [26][26][26]int) string {
	password := make([]byte, length)
	rands := getRandomFloat32Array(length)
	sumFreq := sigma
	randNum := int64(rands[0] * float32(sumFreq))
	sum := 0

init:
	for c1 := 0; c1 < 26; c1++ {
		for c2 := 0; c2 < 26; c2++ {
			for c3 := 0; c3 < 26; c3++ {
				sum += tris[c1][c2][c3]
				if int64(sum) > randNum {
					password[0] = byte('a' + c1)
					password[1] = byte('a' + c2)
					password[2] = byte('a' + c3)
					break init
				}
			}
		}
	}

	// do a random walk
	nchar := 3 // we have three chars so far
	for nchar < length {
		c1 := password[nchar-2] - 'a' // take the last 2 chars
		c2 := password[nchar-1] - 'a' // .. and find the next one
		sumFreq := 0
		for c3 := 0; c3 < 26; c3++ {
			sumFreq += tris[c1][c2][c3]
			// note tha sum < duo[c1][c2] because duos counts all
			// digraphs, not just those in a trigraph. we want sum.
		}
		if sumFreq == 0 { // if there is no possible extension..
			break // break while nchar loop and print what we have
		}
		// choose a continuation
		randNum := int64(rands[nchar] * float32(sumFreq)) // weight by sum of frequencies for row

		sum := 0
		for c3 := 0; c3 < 26; c3++ {
			sum += tris[c1][c2][c3]
			if int64(sum) > randNum%int64(sumFreq) {
				password[nchar] = byte('a' + c3)
				nchar++
				break
			}
		}
	}
	return string(password[:])
}
