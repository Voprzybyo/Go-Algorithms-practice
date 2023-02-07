package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func findUpper(inputName string) []int {
	var trimIndex []int
	for i, v := range inputName {
		if unicode.IsUpper(v) {
			trimIndex = append(trimIndex, i)
		}
	}

	return trimIndex
}

func findSpaces(inputName string) []int {
	var trimIndex []int
	for i, v := range inputName {
		if unicode.IsSpace(v) {
			trimIndex = append(trimIndex, i)
		}
	}

	return trimIndex
}

func removeSpaces(inputName string, trimIndex []int) string {

	for i := range trimIndex {
		inputName = inputName[:trimIndex[i]] + inputName[trimIndex[i]+1:]
		if i < len(trimIndex)-1 {
			trimIndex[i+1] -= i + 1
		}
	}

	return inputName
}

func addSpaces(inputName string, trimIndex []int) string {

	for i := range trimIndex {
		inputName = inputName[:trimIndex[i]] + " " + inputName[trimIndex[i]:]
		if i < len(trimIndex)-1 {
			trimIndex[i+1] += i + 1
		}
	}

	return inputName
}

func nameToUpper(name string, trimIndex []int) string {

	temp := []rune(name)
	for i := range trimIndex {
		if name[trimIndex[i]] >= 'a' && name[trimIndex[i]] < 'z' {
			temp[trimIndex[i]] = temp[trimIndex[i]] - 32
		}
	}
	name = string(temp)

	return name
}

func splitCamelCase(inputData []string) string {

	conversionType := inputData[0]
	inputName := inputData[1]
	var trimIndex []int
	var outputName string

	switch conversionType {
	case "M":
		trimIndex = findUpper(inputName)
		outputName = addSpaces(inputName, trimIndex)
		outputName = outputName[:len(outputName)-2] // Remove "()"
		outputName = strings.ToLower(outputName)

	case "C":
		for i, v := range inputName {
			if i == 0 {
				continue
			}
			if unicode.IsUpper(v) {
				trimIndex = append(trimIndex, i)
			}
		}

		outputName = addSpaces(inputName, trimIndex)
		outputName = strings.ToLower(outputName)

	case "V":
		trimIndex = findUpper(inputName)
		outputName = addSpaces(inputName, trimIndex)
		outputName = strings.ToLower(outputName)
	}

	return outputName
}

func combineCamelCase(inputData []string) string {

	conversionType := inputData[0]
	inputName := inputData[1]
	var trimIndex []int
	var outputName string

	switch conversionType {
	case "M":
		trimIndex = findSpaces(inputName)
		outputName = removeSpaces(inputName, trimIndex)
		outputName = nameToUpper(outputName, trimIndex)
		outputName = outputName + "()"

	case "C":
		trimIndex = findSpaces(inputName)
		outputName = removeSpaces(inputName, trimIndex)
		outputName = nameToUpper(outputName, trimIndex)

		// Change first letter to uppercase
		temp := []rune(outputName)
		temp[0] = temp[0] - 32
		outputName = string(temp)

	case "V":
		trimIndex = findSpaces(inputName)
		outputName = removeSpaces(inputName, trimIndex)
		outputName = nameToUpper(outputName, trimIndex)
	}

	return outputName
}

func formatCamelCase(s string) string {

	var retString string
	inputData := strings.Split(s, ";")
	command := inputData[0]

	switch command {
	case "S":
		retString = splitCamelCase(append(inputData[:0], inputData[1:]...))
	case "C":
		retString = combineCamelCase(append(inputData[:0], inputData[1:]...))
	default:
		fmt.Println("Wrong command!")
	}

	return retString
}

func main() {

	var texts []string
	readFile, err := os.Open("inputData.txt")
	var ret string

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		message := fileScanner.Text()
		if err == nil {
			message = strings.Trim(message, "\n\r")
			texts = append(texts, message)
		} else {
			texts = append(texts, message)
			break
		}
	}

	readFile.Close()

	for _, text := range texts {
		ret = formatCamelCase(text)
		fmt.Println(ret)
	}
}
