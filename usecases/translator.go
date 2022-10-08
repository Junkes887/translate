package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/translate"
	"github.com/Junkes887/translate/model"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func GetTranslate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	text := r.FormValue("text")
	typeLang := r.FormValue("typeLang")
	doResponseTranslate(w, DoTranslate(text, typeLang))
}

func doResponseTranslate(w http.ResponseWriter, word model.Word) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(word)
}

func DoTranslate(text string, typeLang string) model.Word {
	ctx := context.Background()

	// lang, err := language.Parse("pt-br")
	// lang, err := language.Parse("en")
	lang, err := language.Parse(typeLang)
	if err != nil {
		fmt.Println(err)
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		fmt.Println(err)
	}

	translations, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		fmt.Println(err)
	}

	return model.Word{
		Text:           text,
		TranslatedText: translations[0].Text,
	}
}

func DoTranslateList(texts []string, typeLang string) model.Words {
	ctx := context.Background()

	// lang, err := language.Parse("pt-br")
	// lang, err := language.Parse("en")
	lang, err := language.Parse(typeLang)
	if err != nil {
		fmt.Println(err)
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		fmt.Println(err)
	}

	translations, err := client.Translate(ctx, texts, lang, nil)
	if err != nil {
		fmt.Println(err)
	}

	return model.Words{
		Texts:           texts,
		TranslatedTexts: translations,
	}
}
