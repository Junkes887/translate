package main

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/Junkes887/translate/html"
	"github.com/Junkes887/translate/request"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

const URL string = "https://www.google.com/search?q="

func main() {
	// translateText("Hello, world!")
	findResults("teste de velocidade", "0")
}

func findResults(query string, start string) {
	resp := request.Request(makeUrl(query, start))

	html.ManipulateHTML(resp.Body)
}

func makeUrl(query string, start string) string {
	return fmt.Sprintf("%s%s&%s", URL, strings.ReplaceAll(query, " ", "+"), start)
}

func translateText(text string) {
	ctx := context.Background()

	lang, err := language.Parse("pt-br")
	if err != nil {
		fmt.Println(err)
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey("AIzaSyC7UWk4UoMLmG2ZIdYv7des9hfHyXNsd2g"))
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
