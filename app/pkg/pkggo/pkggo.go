package pkggo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gopherzz/gg/app/pkg/pkggo/models"
)

const urlPattern = "https://pkg.go.dev/search?q=%s"

func FindPackages(query string) ([]models.GoPackage, error) {
	res, err := http.Get(fmt.Sprintf(urlPattern, query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	packages := make([]models.GoPackage, 0)

	doc.Find("div.SearchSnippet").Each(func(i int, s *goquery.Selection) {
		url := s.Find("span.SearchSnippet-header-path").Text()
		splittedUrl := strings.Split(url[1:len(url)-1], "/")
		title := splittedUrl[len(splittedUrl)-1]

		shortDesc := strings.TrimSpace(s.Find("p.SearchSnippet-synopsis").Text())

		packages = append(packages, models.GoPackage{
			Name:      title,
			Url:       url,
			ShortDesc: shortDesc,
		})
	})

	return packages, nil
}
