package scrapper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fedeya/lol-save-builds/champeon"
	"github.com/gocolly/colly/v2"
)

var itemsCount int
var listItemsCount int = -1

// GetChampItems get the recomended champion items
func (s *Scrapper) GetChampItems(e *colly.HTMLElement) {
	s.wg.Add(1)
	defer s.wg.Done()
	e.ForEach("li.champion-stats__list__item", func(i int, e *colly.HTMLElement) {
		if itemsCount >= 7 {
			if i == 0 {
				var arr []champeon.Item
				listItemsCount++
				s.Champ.Items = append(s.Champ.Items, arr)
			}

			t := e.Attr("title")
			src := e.ChildAttr("img", "src")
			r := strings.NewReader(t)
			doc, _ := goquery.NewDocumentFromReader(r)
			item := doc.Find("b").Text()
			s.Champ.Items[listItemsCount] = append(s.Champ.Items[listItemsCount], champeon.Item{
				Name:     item,
				ImageURL: src,
			})
		} else {
			itemsCount++
		}
	})
}
