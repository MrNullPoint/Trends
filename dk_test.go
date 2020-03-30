package Trends

import (
	"net/http"
	"testing"
)

func TestDKResponse(t *testing.T) {
	tr := NewTrend()

	keywords, resp, err := tr.DailyKeywordSearch()

	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.Status)
	}

	for _, k := range keywords {
		t.Log(k.Marshal())
	}
}
