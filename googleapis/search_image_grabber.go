package googleapis

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

const GooglePageResultsSize = 10

type IGoogleApiService interface {
	DownloadImages(query string, page int) ([]MemImage, error)
}

type GoogleApiService struct {
	apiKey         string
	searchEngineID string
}

func New(apiKey, searchEngineID string) IGoogleApiService {
	return GoogleApiService{
		apiKey:         apiKey,
		searchEngineID: searchEngineID,
	}
}

func (g GoogleApiService) DownloadImages(query string, page int) ([]MemImage, error) {
	log.Printf("start downloading image with query %s, %d page \n", query, page)
	links, err := g.doSearch(query, page)
	if err != nil {
		return nil, err
	}

	ch := make(chan MemImage, len(links))
	imgs := make([]MemImage, 0, 0)

	var wg sync.WaitGroup
	for _, link := range links {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- grabImage(link)
		}()
	}
	wg.Wait()
	close(ch)

	for r := range ch {
		imgs = append(imgs, r)
	}
	return imgs, nil
}

func (g GoogleApiService) doSearch(query string, page int) ([]string, error) {
	resp, err := http.Get(sanitizeQueryUrl(query, g.apiKey, g.searchEngineID, page))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results GoogleResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	links := make([]string, 0, len(results.Items))
	for _, item := range results.Items {
		links = append(links, item.Link)
	}
	return links, nil
}

func sanitizeQueryUrl(query, apiKey, searchEngineID string, page int) string {
	u, _ := url.Parse("https://www.googleapis.com/customsearch/v1")
	q := u.Query()
	q.Add("q", url.QueryEscape(query))
	q.Add("searchType", "image")
	q.Add("key", apiKey)
	q.Add("cx", searchEngineID)
	q.Add("start", strconv.Itoa(page*GooglePageResultsSize))
	u.RawQuery = q.Encode()

	return u.String()
}

func grabImage(url string) MemImage {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("err in grab image function: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("err in grab image function while storing response: %v\n", err)
		return nil
	}
	return res
}
