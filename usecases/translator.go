package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/translate"
	"github.com/Junkes887/translate/model"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func GetTranslate(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	doResponseTranslate(w, DoTranslate(text))
}

func doResponseTranslate(w http.ResponseWriter, word model.Word) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(word)
}

func DoTranslate(text string) model.Word {
	ctx := context.Background()

	lang, err := language.Parse("pt-br")
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
