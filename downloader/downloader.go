package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// CoverageURL カバレッジ取得用エンドポイントURL
	coverageURL = "https://api.openbd.jp/v1/coverage"
	// BooksURL 書誌データ取得用エンドポイントURL
	booksURL = "https://api.openbd.jp/v1/get"
)

// DownloadCoverage ISBNのリストをダウンロードする
func DownloadCoverage() ([]byte, error) {
	resp, err := http.Get(coverageURL)
	if err != nil {
		return nil, fmt.Errorf("http response error: %w", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http response read error: %v", err)
	}

	return b, nil
}

// DownloadBookInfo 書誌データのダウンローダー
func DownloadBookInfo(list []string) ([]byte, error) {

	data := url.Values{"isbn": {strings.Join(list, ",")}}
	resp, err := http.PostForm(booksURL, data)
	if err != nil {
		return nil, fmt.Errorf("http response error: %w", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http response read error: %v", err)
	}

	return b, nil
}
