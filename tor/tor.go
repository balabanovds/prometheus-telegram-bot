package tor

import (
	"net/http"
	"net/url"

	"github.com/balabanovds/prometheus-telegram-bot/util"
)

// GetClient returns *http.Client depending of config.yml tor.http-proxy
func GetClient() *http.Client {
	client := &http.Client{}

	proxyURL, err := url.Parse(util.Cfg.Tor.HTTPProxy)
	if err != nil {
		return client
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client.Transport = transport

	return client
}
