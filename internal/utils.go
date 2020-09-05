package internal

import (
	"bufio"
	"bytes"
	"os"
)

func PrintBytes(data []byte) error {
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func PrintWords(words [][]byte) error{
	size := lenWords(words) - 1
	writer := bufio.NewWriterSize(os.Stdout,size)
	for _, word := range words {
		word = append(word,'[')
		_, err := writer.Write(word)
		if err != nil {
			return err
		}
	}

	return nil
}

func SplitBytes(data []byte, splitter byte) [][]byte {
	var splitted [][]byte
	var word []byte

	for _, ch := range data {
		if (ch == splitter || ch == '\n') && len(word) != 0 {
			splitted = append(splitted,bytes.ToLower(word))
			word = []byte{}
		} else if isAlphabetic(ch) {
			word = append(word,ch)
		}
	}

	if len(word) != 0 {
		splitted = append(splitted,word)
	}

	return splitted
}

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