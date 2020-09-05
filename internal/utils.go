package internal

func isByteMatrixesEqual(a [][]byte, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func isAlphabetic(ch byte) bool{
	if (ch > 64 && ch < 91) || (ch > 96 && ch < 123) {
		return true
	}

	return false
}


func lenWords(words[][]byte) int {
	var lenW int

	for _, word := range words {
		lenW += len(word) + 1
	}

	return lenW
}