package siml

import (
	"fmt"
	"strings"

	"crypto/tls"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Get(domain string) ([]string, error) {
	var list []string

	url := fmt.Sprintf(endpoint, domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return list, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", accept)
	req.Header.Set("Accept-Language", acceptLang)

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return list, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return list, err
	}

	doc.Find(jsSelector).Each(func(i int, s *goquery.Selection) {
		h := strings.TrimSpace(s.Text())
		if h != "" {
			list = append(list, h)
		}
	})

	return list, nil
}
