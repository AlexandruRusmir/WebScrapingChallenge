package services

import (
	"api/util"
)

var positiveWords = map[string]bool{
	"good":          true,
	"happy":         true,
	"excellent":     true,
	"positive":      true,
	"superb":        true,
	"beneficial":    true,
	"bright":        true,
	"delightful":    true,
	"fantastic":     true,
	"favorable":     true,
	"impressive":    true,
	"pleasing":      true,
	"remarkable":    true,
	"splendid":      true,
	"valuable":      true,
	"joy":           true,
	"victory":       true,
	"beautiful":     true,
	"best":          true,
	"better":        true,
	"winning":       true,
	"blessing":      true,
	"bliss":         true,
	"celebration":   true,
	"peace":         true,
	"praise":        true,
	"enjoy":         true,
	"glad":          true,
	"healthy":       true,
	"perfect":       true,
	"progress":      true,
	"prosperous":    true,
	"joys":          true,
	"successful":    true,
	"worthy":        true,
	"profitable":    true,
	"superior":      true,
	"upbeat":        true,
	"vibrant":       true,
	"love":          true,
	"amazing":       true,
	"great":         true,
	"win":           true,
	"achievement":   true,
	"reliability":   true,
	"benefit":       true,
	"collaboration": true,
	"commitment":    true,
	"confidence":    true,
}

var negativeWords = map[string]bool{
	"bad":           true,
	"sad":           true,
	"terrible":      true,
	"negative":      true,
	"adverse":       true,
	"awful":         true,
	"disagreeable":  true,
	"dreadful":      true,
	"evil":          true,
	"harmful":       true,
	"poor":          true,
	"unacceptable":  true,
	"regret":        true,
	"unfavorable":   true,
	"fail":          true,
	"loss":          true,
	"depressing":    true,
	"harsh":         true,
	"horrible":      true,
	"inferior":      true,
	"miserable":     true,
	"unpleasant":    true,
	"upset":         true,
	"worthless":     true,
	"deficit":       true,
	"disappoint":    true,
	"frustrate":     true,
	"lament":        true,
	"misery":        true,
	"pitiful":       true,
	"sorrow":        true,
	"trouble":       true,
	"annoying":      true,
	"tragic":        true,
	"unhappy":       true,
	"woe":           true,
	"risky":         true,
	"ugly":          true,
	"annoy":         true,
	"anxiety":       true,
	"boring":        true,
	"complicated":   true,
	"concern":       true,
	"disgusting":    true,
	"distress":      true,
	"disturb":       true,
	"disappointing": true,
	"unfortunate":   true,
	"unfair":        true,
}

func AnalyzeSentimentForURL(url string) (string, error) {
	data, err := ScrapeContent(url, nil)
	if err != nil {
		return "", err
	}

	combinedText := util.CombineTextFromContent(data)
	return AnalyzeSentiment(combinedText), nil
}

func AnalyzeSentiment(text string) string {
	tokens := util.Tokenize(text)

	score := 0
	for _, token := range tokens {
		if positiveWords[token] {
			score++
		}
		if negativeWords[token] {
			score--
		}
	}

	if score > 0 {
		return "positive"
	} else if score < 0 {
		return "negative"
	}
	return "neutral"
}
