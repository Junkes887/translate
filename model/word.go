package model

import "cloud.google.com/go/translate"

// Word entity
type Word struct {
	Text           string
	TranslatedText string
}

// Word entity
type Words struct {
	Texts           []string
	TranslatedTexts []translate.Translation
}
