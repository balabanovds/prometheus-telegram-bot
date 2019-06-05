package prometheus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	u "github.com/balabanovds/prometheus-telegram-bot/util"
)

type prometheusAPIResponse struct {
	Data struct {
		Result []struct {
			Metric struct {
				Name       string `json:"__name__"`
				Device     string `json:"device"`
				FsType     string `json:"fstype"`
				Mountpoint string `json:"mountpoint"`
				Instance   string `json:"instance"`
				Backend    string `json:"backend"`
				Server     string `json:"server"`
				Job        string `json:"job"`
				State      string `json:"state"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
		ResultType string `json:"resultType"`
	} `json:"data"`
	Status string `json:"status"`
}

// Response is final response from API
type Response struct {
	APIResponse *prometheusAPIResponse
	Query       *u.Query
	Err         error
}

// FetchAll requests all queries from config
func FetchAll() []*Response {
	var responses []*Response

	queries := u.Cfg.Prometheus.Queries

	var wg sync.WaitGroup
	wg.Add(len(queries))

	for _, q := range queries {
		go func(q u.Query) {
			defer wg.Done()

			var sb strings.Builder
			sb.WriteString(u.Cfg.Prometheus.API)
			sb.WriteString(`query?query=`)
			sb.WriteString(q.Query)
			res, err := getAPIResponse(sb.String())

			resp := Response{
				Query: &q,
			}

			if err != nil {
				resp.Err = err
			} else {
				resp.APIResponse = res
			}

			responses = append(responses, &resp)

		}(q)
	}
	wg.Wait()
	return responses
}

func getAPIResponse(url string) (*prometheusAPIResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status %v", res.StatusCode)
	}

	var resp prometheusAPIResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
