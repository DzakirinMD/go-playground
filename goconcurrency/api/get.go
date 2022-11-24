package api

import (
	"encoding/json"
	"fmt"
	"go_playground/goconcurrency/structs"
	"net/http"
	"strings"
	"time"
)

const Url = "https://xkcd.com"

/*
fetch url
custom HTTP client
timeout in 5 second
*/
func Fetch(n int) (*structs.Result, error) {

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	// concatenate strings to get url; ex: https://xkcd.com/571/info.0.json
	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}

	var data structs.Result

	// error from web service, empty struct to avoid disruption of process
	if resp.StatusCode != http.StatusOK {
		data = structs.Result{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}

	resp.Body.Close()

	return &data, nil
}
