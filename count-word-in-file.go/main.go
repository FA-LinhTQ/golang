package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func countWordInFile(ch chan int, filePath, keyword string) {
	var numberOfOcc int
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		ch <- 0
		return
	}
	numberOfOcc = strings.Count(string(fileContent), keyword)
	ch <- numberOfOcc
	defer close(ch)
}

func main() {
	countFirsrCh := make(chan int)
	countSecondCh := make(chan int)

	go countWordInFile(countFirsrCh, "1.txt", "a")
	go countWordInFile(countSecondCh, "2.txt", "a")
	fmt.Println("Total: ", <-countFirsrCh+<-countSecondCh)

}
