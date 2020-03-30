package Trends

import (
	"net/http"
	"strings"
)

const GoogleDailyRss = "https://trends.google.com/trends/trendingsearches/daily/rss"
const UserAgent = "mrnullpoint-go-trends 0.1"

type Trend struct {
	client   *http.Client
	request  *http.Request
	response *http.Response
	geo      string
	t        int
}

func NewTrend() *Trend {
	tr := new(Trend)
	tr.client = new(http.Client)
	return tr
}

// @function: set your own http client option, eg: http client with proxy
func (tr *Trend) Client(c *http.Client) {
	tr.client = c
}

// @function: set your own geo code, eg: us or US
func (tr *Trend) Geo(geo string) {
	tr.geo = strings.ToUpper(geo)
}

// @function: do http request to fetch daily keywords
func (tr *Trend) DailyKeywordSearch() ([]*DailyKeyword, *http.Response, error) {
	var err error

	if err = tr.buildDKReq(); err != nil {
		return nil, nil, err
	}

	if tr.response, err = tr.client.Do(tr.request); err != nil {
		return nil, tr.response, err
	} else {
		return tr.parseDKResp()
	}
}
