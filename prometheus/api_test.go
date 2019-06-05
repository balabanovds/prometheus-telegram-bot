package prometheus

import (
	"testing"

	"github.com/balabanovds/prometheus-telegram-bot/util"
)

func TestFetchAll(t *testing.T) {
	util.Init()
	FetchAll()
}
