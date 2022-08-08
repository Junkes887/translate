package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func main() {
	translateText("Hello, world!")
}

func translateText(text string) {
	ctx := context.Background()

	lang, err := language.Parse("ru")
	if err != nil {
		fmt.Println(err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}

	translations, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)
}
