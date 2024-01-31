package getembedurl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Youtube struct {
	Html string `json:"html"`
}

func GetEmbedUrl(url string) string {
	urlStr := fmt.Sprintf("https://www.youtube.com/oembed?url=%s", url)
	resp, err := http.Get(urlStr)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var data Youtube
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(data.Html))

	var src string
	doc.Find("iframe").Each(func(i int, s *goquery.Selection) {
		var exists bool
		src, exists = s.Attr("src")
		if !exists {
			fmt.Fprintln(os.Stderr, "src attribute not found")
			os.Exit(1)
		}
	})
	return src
}
