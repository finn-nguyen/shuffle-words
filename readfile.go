package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"math/rand"
	"time"
)

func shuffle(word string) string {
	array := strings.Split(word, "")
	length := len(array)
	newArr := array[:]
	for i, _ := range newArr {
		newIndex := rand.Intn(length)
		temp := newArr[i]
		newArr[i] = newArr[newIndex]
		newArr[newIndex] = temp
	}
	return strings.Join(newArr, "")
}

	
func check(e error) {
	if e != nil {
			panic(e)
	}
}

func main() {
	buffer, err := ioutil.ReadFile("/Users/imac27/Desktop/test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		text := string(buffer)
		vocabulaires := strings.Split(strings.Replace(text, "\r\n", "\n", -1), "\n")
		newWords := vocabulaires[:]
		rand.Seed(time.Now().UnixNano())
		for index := range newWords {
			newWords[index] = shuffle(newWords[index])
		}
		text = strings.Join(newWords, "\r\n")
		fmt.Println(text)
		buffer = []byte(text)
		newFile, err := os.Create("/Users/imac27/Desktop/test_result.txt")
		check(err)
		defer newFile.Close()
		_, err = newFile.Write(buffer)
		check(err)
	}
}