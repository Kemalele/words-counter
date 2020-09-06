package internal

import "bytes"

func itob(num int) []byte {
	var result []byte
	const asciiIntegerStartingIndex = 48
	for num != 0 {
		b := byte(num % 10 + asciiIntegerStartingIndex)
		result = append([]byte{b}, result...)
		num /= 10
	}


	return result
}

func isAlphabetic(ch byte) bool{
	if (ch > 64 && ch < 91) || (ch > 96 && ch < 123) {
		return true
	}

	return false
}

func wordsSame(a, b []word) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !bytes.Equal(a[i].Data,b[i].Data) || a[i].Count != b[i].Count {
			return false
		}
	}

	return true
}