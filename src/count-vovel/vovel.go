package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func takeInputFromUserScanf(prompt string) (string, error) {
	fmt.Print(prompt)
	var userInput string
	if _, err := fmt.Scanf("%s", &userInput); err != nil {
		return "", err
	}
	return userInput, nil
}

func takeInputFromUser(prompt string) (string, error) {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	return line, nil
}

func CreateInputPrompt(rd io.Reader) func(string) (string, error) {
	in := bufio.NewReader(rd)
	return func(prompt string) (string, error) {
		fmt.Print(prompt)
		line, err := in.ReadString('\n')
		if err != nil {
			return "", err
		}
		return line, nil
	}
}

func isVovel(char int32) bool {
	return char == 'a' || char == 'e' || char == 'u' || char == 'i' || char == 'o'
}

func main() {
	inputPrompt := CreateInputPrompt(os.Stdin)
	message, _ := inputPrompt("Enter a text-> ")
	lowerMessage := strings.ToLower(message)
	vovel, consonant := 0, 0
	for _, char := range lowerMessage {
		if char >= 'a' && char <= 'z' {
			if isVovel(char) {
				vovel++
			} else {
				consonant++
			}
		}
	}
	fmt.Printf("vovel: %d, consonant: %d\n", vovel, consonant)
}
