package repository

import (
	"encoding/json"
	"fmt"
	"github.com/emiliano080591/concurrency/project/model"
	"net/http"
	"strings"
	"time"
)

const (
	Url    = "https://xkcd.com"
	METHOD = "GET"
)

func Fetch(n int) (result *model.Result, err error) {
	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	// concatenate strings to get url; ex: https://xkcd.com/571/info.0.json
	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")

	req, err := http.NewRequest(METHOD, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}

	var data model.Result

	// error from web service, empty struct to avoid disruption of process
	if resp.StatusCode != http.StatusOK {
		data = model.Result{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}
	resp.Body.Close()

	return &data, nil
}
