package parser

import "encoding/json"

type coverage struct {
}

// ParserCoverageJSON /v1//coverageから取得したJSONをパースし、isbnのリストを返す
func ParserCoverageJSON(data []byte) ([]string, error) {
	var j []string
	err := json.Unmarshal(data, &j)

	if err != nil {
		return nil, err
	}

	return j, nil
}

// ParseBookInfoJson /v1/getから取得したJSONをパースし、書誌データを返す
func ParseBookInfoJSON(data []byte) (interface{}, error) {
	var j interface{}
	err := json.Unmarshal(data, &j)

	if err != nil {
		return nil, err
	}

	return j, nil
}
