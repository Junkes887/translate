package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	htmlHandler "github.com/Junkes887/translate/handlers"
	"github.com/Junkes887/translate/model"
	"github.com/Junkes887/translate/request"
	"github.com/julienschmidt/httprouter"
)

const URL_TEMPLATE string = "https://www.google.com/search?q=%s&start=%s&gl=us&gws_rd=ssl"

func GetTranslateAndSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//pageListMock := []model.Page{}
	//for i := 0; i < 10; i++ {
	//	pageListMock = append(pageListMock, model.Page{Description: "Descrição", Title: "Titúlo", Link: "http://www.google.com"})
	//}
	query := r.FormValue("query")
	start := r.FormValue("start")
	response := doRequest(query, start)
	pageList := htmlHandler.ManipulateHTML(response.Body)
	doResponseTranslateAndSearch(w, pageList)
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
