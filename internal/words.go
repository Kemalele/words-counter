package internal

import (
	"bufio"
	"bytes"
	"os"
)
const recentLimit = 20

type word struct {
	Data []byte
	Count int
}

type Words struct {
	Body 		     []word
	mostRecents      []word
}

func (words *Words) Append(w word) {
	// initializing mostRecents array
	if len(words.mostRecents) < recentLimit {
		words.mostRecents = make([]word,recentLimit)
	}

	if i := words.exists(w); i != -1 {
		words.Body[i].Count += 1
	} else {
		words.Body = append(words.Body, w)
	}

	words.mostRecent(w)
}

func (words *Words) exists(w word) int {
	if len(words.Body) == 1 && bytes.Equal(words.Body[0].Data,w.Data){
		return 0
	}

	// front, back search
	for i, j := 0, len(words.Body) -1; i < j; i,j = i+1, j-1 {
		if bytes.Equal(words.Body[i].Data, w.Data) {
			return i
		} else if bytes.Equal(words.Body[j].Data, w.Data) {
			return j
		}
	}

	return -1
}

func (words *Words) mostRecent(w word) bool {
	for i, word := range words.mostRecents {
		if w.Count > word.Count {
			words.mostRecents[i] = w
			return true
		}
	}
	return false
}

func Solve(file []byte) error{
	words := split(file)
	err := printWords(words)
	if err != nil {
		return err
	}

	return nil
}


func split(data []byte) Words {
	var words Words
	var text []byte

	for _, ch := range data {
		if !isAlphabetic(ch) && len(text) != 0 {
			w := word{Data:bytes.ToLower(text)}
			words.Append(w)
			text = []byte{}
		} else if isAlphabetic(ch) {
			text = append(text,ch)
		}
	}

	if len(text) != 0 {
		w := word{Data:bytes.ToLower(text)}
		words.Append(w)
	}

	return words
}

func printWords(words Words) error{
	size := len(words.Body) - 1
	writer := bufio.NewWriterSize(os.Stdout,size)
	for _, word := range words.Body {
		word.Data = append(word.Data,'[')
		_, err := writer.Write(word.Data)
		if err != nil {
			return err
		}
	}

	return nil
}


