package siml

import (
	"errors"
	"fmt"
	"strings"

	"crypto/tls"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Get returns a slice of similar websites for the given domain,
// as scraped from https://www.sitelike.org/similar/{domain}/.
//
// It returns an error if the request fails or if the response cannot be parsed.
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

// Gets same as Get but it accept multiple domain inputs.
func Gets(domains ...string) ([]string, error) {
	var list []string

	switch len(domains) {
	case 0:
		return list, errors.New(errNoDomainInput)
	case 1:
		return Get(domains[0])
	}

	errCh := make(chan error, len(domains))

	for _, domain := range domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			similar, err := Get(domain)
			if err != nil {
				errCh <- err
				return
			}
			list = append(list, similar...)
		}(domain)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}
