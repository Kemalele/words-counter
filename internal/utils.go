package internal

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

func bytesToLower(input []byte) []byte{
	var res []byte
	const lowerCaseDifference = 32
	for _, b := range input {
		if b > 64 && b < 91 {
			upperCase := b + lowerCaseDifference
			res = append(res, upperCase)
		} else {
			res = append(res,b)
		}
	}

	return res
}

func bytesJoin(s [][]byte,sep []byte) []byte{
	var result []byte
	for i, bytes := range s {
		result = append(result,bytes...)
		// Not adding separator to last element
		if i != len(s) -1 {
			result = append(result,sep...)
		}
	}

	return result
}

func isAlphabetic(ch byte) bool{
	if (ch > 64 && ch < 91) || (ch > 96 && ch < 123) {
		return true
	}

	return false
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func wordsSame(a, b []word) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !bytesEqual(a[i].Data,b[i].Data) || a[i].Count != b[i].Count {
			return false
		}
	}

	return true
}