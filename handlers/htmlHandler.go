package handlers

import (
	"io"
	"regexp"
	"strings"

	"github.com/Junkes887/translate/artifacts"
	"github.com/Junkes887/translate/model"
	"github.com/PuerkitoBio/goquery"
)

func ManipulateHTML(res io.ReadCloser) model.Page {
	doc, err := goquery.NewDocumentFromReader(res)
	artifacts.HandlerError(err)

	var sites []model.Site

	var title, description, link, total string

	doc.Find("div#result-stats").Each(func(i int, s *goquery.Selection) {
		texto := s.Get(0).FirstChild.Data
		pattern := regexp.MustCompile(`[0-9-,]+`)
		firstMatchIndex := pattern.FindAllString(texto, -1)

		for _, element := range firstMatchIndex {
			total = strings.ReplaceAll(element, ",", ".")
		}

	})

	doc.Find(".MjjYud").Each(func(i int, selectionFather *goquery.Selection) {
		class := selectionFather.Get(0).Parent.Attr[0]

		if class.Val != "v7W49e" {
			return
		}

		selectionFather.Find(".yuRUbf > a > h3").Each(func(i int, s *goquery.Selection) {
			title = s.Get(0).FirstChild.Data
		})
		selectionFather.Find(".VwiC3b.yXK7lf.MUxGbd.yDYNvb.lyLwlc.lEBKkf").Each(func(i int, s *goquery.Selection) {

			first := s.Get(0).FirstChild

			if first.Data == "span" {
				s.Find("span").Each(func(i int, child *goquery.Selection) {
					m := regexp.MustCompile("<[^>]*>")
					outer, _ := goquery.OuterHtml(child)
					description = m.ReplaceAllString(outer, "")
				})
			} else {
				description = first.Data
			}
		})
		if description == "" {
			selectionFather.Find(".VwiC3b.yXK7lf.MUxGbd.yDYNvb.lyLwlc").Each(func(i int, s *goquery.Selection) {

				first := s.Get(0).FirstChild

				if first.Data == "span" {
					s.Find("span").Each(func(i int, child *goquery.Selection) {
						m := regexp.MustCompile("<[^>]*>")
						outer, _ := goquery.OuterHtml(child)
						description = m.ReplaceAllString(outer, "")
					})
				} else {
					description = first.Data
				}
			})
		}

		selectionFather.Find(".yuRUbf > a").Each(func(i int, s *goquery.Selection) {
			link, _ = s.Attr("href")
		})

		if title == "" || description == "" {
			return
		}
		sites = append(sites, model.Site{
			OriginalTitle:       title,
			OriginalDescription: description,
			Link:                link,
		})

		description = ""
	})

	return model.Page{
		Sites: sites,
		Total: total,
	}
}
