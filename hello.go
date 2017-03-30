package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	// "time"
	"math/rand"
	"strings"
)

// func randomPos(min, max int) (int) {
// 	result := rand.Intn(max - min) + min
// 	return result
// }

func main() {
	// TODO: input and output paths should be command line arguments
	inputFilePath := "in.txt"

	// seed with non-predictable value
	// rand.Seed(time.Now().UTC().UnixNano())

	contentBytes, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	content := string(contentBytes)

	result := shuffleSentenceParts(content)
	//result := shuffleWords(content)

	fmt.Println(result)
}

type wordFormatter func(string) string

func formatSentencePart(part string) string {

	firstLetter := strings.ToUpper(part[:1])
	result := firstLetter + part[1:] + "\n"

	return result
}

func formatWordPart(part string) string {

	result := part + " "

	return result
}

func shuffleSentenceParts(input string) string {

	regexSplit := ", |\\. | and | but | or | which |\\n"
	result := shuffleParts(input, regexSplit, formatSentencePart)
	
	return result
}

func shuffleWords(input string) string {

	regexSplit := "\\s"
	result := shuffleParts(input, regexSplit, formatWordPart)
	
	return result
}

func shuffleParts(input, regexSplit string, formatter wordFormatter) string {

	result := ""

	regex := regexp.MustCompile(regexSplit)
	parts := regex.Split(input, -1)
	
	indexes := rand.Perm(len(parts))

	for _, index := range indexes {
		if len(parts[index]) > 0 {
			result += formatter(parts[index])
		}
	}

	return result
	
}

