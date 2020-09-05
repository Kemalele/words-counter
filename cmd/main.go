package main
import (
	//"fmt"
	//"io"
	"io/ioutil"
	"log"
	internal "../internal"
)
func main() {
	f, err :=  ioutil.ReadFile("../static/mobydick.txt")
	if err != nil {
		log.Panic(err)
	}

	err = internal.Solve(f)

	if err != nil {
		log.Panic(err)
	}

}