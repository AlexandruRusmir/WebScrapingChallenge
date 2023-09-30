package util

import (
	"regexp"
	"strings"
)

func Tokenize(text string) []string {
	reg := regexp.MustCompile("[a-zA-Z0-9_]+")
	return reg.FindAllString(strings.ToLower(text), -1)
}

func GetWordCount(text string) int {
	tokens := Tokenize(text)
	return len(tokens)
}

func CombineTextFromContent(data map[string]interface{}) string {
	var combinedText string
	for _, content := range data {
		switch content := content.(type) {
		case []string:
			combinedText += strings.Join(content, " ")
		case []map[string]string:
			for _, item := range content {
				if text, exists := item["text"]; exists {
					combinedText += text + " "
				}
			}
		}
	}

	return combinedText
}
