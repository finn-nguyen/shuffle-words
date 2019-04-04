package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func shuffle(word string) string {
	newWord := ""
	for _, character := range word {
		if character % 2 == 0 {
			newWord = newWord + string(character)
		} else {
			newWord = string(character) + newWord
		}
	}
	return newWord
}

	
func check(e error) {
	if e != nil {
			panic(e)
	}
}

func main() {
	buffer, err := ioutil.ReadFile("C:\\Users\\nvtuo\\Desktop\\vocabularies.txt")
	fmt.Println(shuffle("123456789"))
	if err != nil {
		fmt.Println(err)
	} else {
		text := string(buffer)
		vocabulaires := strings.Split(strings.Replace(text, "\r\n", "\n", -1), "\n")
		newWords := vocabulaires[:]
		for index := range newWords {
			newWords[index] = shuffle(newWords[index])
		}
		text = strings.Join(newWords, "\r\n")
		fmt.Println(text)
		buffer = []byte(text)
		newFile, err := os.Create("C:\\Users\\nvtuo\\Desktop\\vocabularies_shuffle.txt")
		check(err)
		defer newFile.Close()
		_, err = newFile.Write(buffer)
		check(err)
	}
}