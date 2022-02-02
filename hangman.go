package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Hangman struct {
	GuessWord        string
	HiddenWord       []rune
	DisplayWord      string
	InputUser        string
	LastPropositions []string
	Trying           int
}

var WORDS = []string{}

func HangmanInit() {
	rand.Seed(time.Now().UTC().UnixNano())
	hangman = Hangman{
		GuessWord:        strings.ToUpper(WORDS[rand.Intn(len(WORDS))]),
		HiddenWord:       []rune{},
		Trying:           0,
		LastPropositions: []string{},
	}
	for range hangman.GuessWord {
		hangman.HiddenWord = append(hangman.HiddenWord, '_')
	}
	resultHangman()
}
func TryInput(input string) {
	if len(input) == 0 {
		return
	} else if len(input) == 1 {
		gAwnser := false
		for i := 0; i < len(hangman.GuessWord); i++ {
			if string(hangman.GuessWord[i]) == strings.ToUpper(input) {
				gAwnser = true
				hangman.HiddenWord[i] = rune(strings.ToUpper(input)[0])
			}
		}
		if !gAwnser {
			hangman.Trying += 1
		}
	} else {
		if strings.ToUpper(input) == hangman.GuessWord {
			hangman.HiddenWord = []rune(hangman.GuessWord)
		}
	}
	alIn := false
	for _, proposition := range hangman.LastPropositions {
		if proposition == strings.ToUpper(input) {
			alIn = true
			break
		}
	}
	if !alIn {
		hangman.LastPropositions = append(hangman.LastPropositions, strings.ToUpper(input))
	}

	resultHangman()
}
func resultHangman() {
	hangman.DisplayWord = string(hangman.HiddenWord)
	if string(hangman.HiddenWord) == hangman.GuessWord {
		hangman.DisplayWord = "You win !"
	} else if hangman.Trying >= 10 {
		hangman.DisplayWord = "You loose..."
	}
}
func LoadingWord(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	word := ""
	for _, char := range string(content) {
		if char == '\n' {
			WORDS = append(WORDS, strings.ToUpper(word[:len(word)-1]))
			word = ""
		} else {
			word += string(char)
		}
	}
}
