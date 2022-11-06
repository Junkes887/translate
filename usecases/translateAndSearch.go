package usecases

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"

	htmlHandler "github.com/Junkes887/translate/handlers"
	"github.com/Junkes887/translate/model"
	"github.com/Junkes887/translate/request"
	"github.com/julienschmidt/httprouter"
)

const URL_TEMPLATE string = "https://www.google.com/search?q=%s&start=%s&gl=us&gws_rd=ssl"
const ENGLISH = "en"
const PORTUGUES = "pt-br"

var cache = make(map[string][]model.Page)

func GetTranslateAndSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.FormValue("query")
	start := r.FormValue("start")
	queryTranslated := DoTranslate(query, ENGLISH)
	var returnValue []model.Page
	url := makeUrl(queryTranslated.TranslatedText, start)
	if ret, ok := cache[url]; ok {
		fmt.Println("Requisição cacheada, retornando da memória.")
		returnValue = ret
	} else {
		response := doRequest(url)
		returnValue = htmlHandler.ManipulateHTML(response.Body)
		doFormatHtml(returnValue)
		cache[url] = returnValue
	}
	doResponseTranslateAndSearch(w, returnValue)
}

func doFormatHtml(pageList []model.Page) {
	var texts []string
	if len(pageList) == 0 {
		return
	}
	for _, p := range pageList {
		texts = append(texts, html.UnescapeString(p.OriginalDescription))
		texts = append(texts, html.UnescapeString(p.OriginalTitle))
	}

	textsTranslated := DoTranslateList(texts, PORTUGUES)

	index := 0

	for i := 0; i < len(textsTranslated.Texts); i += 2 {

		pageList[index].Description = html.UnescapeString(textsTranslated.TranslatedTexts[i].Text)
		pageList[index].Title = html.UnescapeString(textsTranslated.TranslatedTexts[i+1].Text)

		index++
	}
}

func doResponseTranslateAndSearch(w http.ResponseWriter, pageList []model.Page) {
	w.Header().Set("Content-type", "application/json;")
	json.NewEncoder(w).Encode(pageList)
}

func doRequest(url string) *http.Response {
	fmt.Println("Relizando request para: " + url)
	return request.Request(url)
}

func makeUrl(query string, start string) string {
	formattedQuery := strings.ReplaceAll(query, " ", "+")
	return fmt.Sprintf(URL_TEMPLATE, formattedQuery, start)
}
