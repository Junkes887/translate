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

func GetTranslateAndSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//pageListMock := []model.Page{}
	//for i := 0; i < 10; i++ {
	//	pageListMock = append(pageListMock, model.Page{Description: "Descrição", Title: "Titúlo", Link: "http://www.google.com"})
	//}
	query := r.FormValue("query")
	queryTranslated := DoTranslate(query, ENGLISH)
	start := r.FormValue("start")
	response := doRequest(queryTranslated.TranslatedText, start)
	pageList := htmlHandler.ManipulateHTML(response.Body)
	doFormatHtml(pageList)
	doResponseTranslateAndSearch(w, pageList)
}

func doFormatHtml(pageList []model.Page) {
	var texts []string
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

func doRequest(query string, start string) *http.Response {
	fmt.Print(makeUrl(query, start))
	return request.Request(makeUrl(query, start))
}

func makeUrl(query string, start string) string {
	formattedQuery := strings.ReplaceAll(query, " ", "+")
	return fmt.Sprintf(URL_TEMPLATE, formattedQuery, start)
}
