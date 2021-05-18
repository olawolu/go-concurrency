package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"sync"
)

var finalString string

func deliverToFinal(letterChannel chan string, wg *sync.WaitGroup) {
	// recieve incoming string from the channel
	letter := <-letterChannel

	// append the recieved string to the final string 
	finalString += letter
	wg.Done()
}

func capitalize(letterChannel chan string, currentLetter string, wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter)
	defer wg.Done()
	letterChannel <- thisLetter
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var letterChannel chan string = make(chan string)

	// open file
	byteContent, err := ioutil.ReadFile("sample_text.txt")
	if err != nil {
		log.Panicln(err)
	}

	length := len(byteContent)
fmt.Println(length)
	for i := 0; i < length; i++ {
		wg.Add(2)
		go capitalize(letterChannel, string(byteContent[i]), &wg)
		go deliverToFinal(letterChannel, &wg)
		wg.Wait()
	}

	fmt.Println(finalString)
}
