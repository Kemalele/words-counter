package internal

import (
	"bufio"
	"os"
)

func Solve(file []byte) error{
	words := split(file)
	err := printWords(words.mostRecents)
	if err != nil {
		return err
	}

	return nil
}


func split(data []byte) Words {
	words := InitWords(20)
	var text []byte

	for _, ch := range data {
		if !isAlphabetic(ch) && len(text) != 0 {
			w := word{Data:bytesToLower(text)}
			words.Append(w)
			text = []byte{}
		} else if isAlphabetic(ch) {
			text = append(text,ch)
		}
	}

	if len(text) != 0 {
		w := word{Data:bytesToLower(text)}
		words.Append(w)
	}

	return words
}

func printWords(words []word) error{
	var length int
	for _,word := range words {
		length += word.Count
	}

	writer := bufio.NewWriterSize(os.Stdout,1)
	for _, word := range words {
		var output []byte
		word.Data = append(word.Data,'\n')
		output = bytesJoin([][]byte{itob(word.Count),word.Data},[]byte(" "))
		_, err := writer.Write(output)
		if err != nil {
			return err
		}
	}

	return nil
}
