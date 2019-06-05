package prometheus

import (
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

type Response struct {
	APIResponse prometheusAPIResponse
	Name        string
	Err         error
}

func GetAllAPIResponses() []Response {
	var responses []Response

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

			res, err := http.Get(sb.String())
			if err != nil {
				responses = append(response
			}

		}(q)
	}

	return responses
}
