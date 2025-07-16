package main

import "strings"

func cleanInput(text string) []string {
	lowText := strings.ToLower(text)
	words := strings.Fields(lowText)
	return words
}
