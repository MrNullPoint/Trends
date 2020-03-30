package Trends

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RssRoot struct {
	XMLName xml.Name   `json:"-" xml:"rss"`
	Channel RssChannel `json:"-" xml:"channel"`
}

type RssChannel struct {
	XMLName xml.Name   `json:"-" xml:"channel"`
	Item    []*RssItem `json:"-" xml:"item"`
}

type RssItem struct {
	XMLName       xml.Name       `json:"-" xml:"item"`
	Title         string         `xml:"title"`
	Traffic       string         `xml:"approx_traffic"`
	Description   string         `xml:"description"`
	Link          string         `xml:"link"`
	PubDate       string         `xml:"pubDate"`
	Picture       string         `xml:"picture"`
	PictureSource string         `xml:"picture_source"`
	NewsItem      []*RssNewsItem `xml:"news_item"`
}

type RssNewsItem struct {
	XMLName xml.Name `json:"-" xml:"news_item"`
	Title   string   `xml:"news_item_title"`
	Snippet string   `xml:"news_item_snippet"`
	Url     string   `xml:"news_item_url"`
	Source  string   `xml:"news_item_source"`
}

type DailyKeyword struct {
	Keyword        string    `json:"keyword"`
	RelatedKeyword []string  `json:"related_keyword"`
	Link           string    `json:"link"`
	Traffic        int64     `json:"traffic"`
	Date           time.Time `json:"date"`
	Picture        string    `json:"picture"`
	PictureSource  string    `json:"picture_source"`
	RelateNews     []*News   `json:"relate_news"`
}

type News struct {
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
	Url     string `json:"url"`
	Source  string `json:"source"`
}

// @function: build daily keyword request
func (tr *Trend) buildDKReq() error {
	var err error

	tr.request, err = http.NewRequest("GET", GoogleDailyRss, nil)

	if err != nil {
		return err
	}

	if tr.geo != "" {
		tr.request.URL.Query().Add("geo", tr.geo)
	}

	tr.request.Header.Set("user-agent", UserAgent)

	return nil
}

// @function: parse daily keyword response
func (tr *Trend) parseDKResp() ([]*DailyKeyword, *http.Response, error) {
	defer tr.response.Body.Close()

	b, _ := ioutil.ReadAll(tr.response.Body)

	var rss RssRoot
	if err := xml.Unmarshal(b, &rss); err != nil {
		return nil, tr.response, err
	}

	var ret []*DailyKeyword
	for _, item := range rss.Channel.Item {
		dk := new(DailyKeyword)

		dk.Keyword = item.Title
		dk.RelatedKeyword = strings.Split(item.Description, ",")
		dk.Date, _ = time.Parse(time.RFC1123Z, item.PubDate)
		dk.Traffic, _ = strconv.ParseInt(strings.ReplaceAll(strings.Trim(item.Traffic, "+"), ",", ""), 10, 64)
		dk.Link = item.Link
		dk.Picture = item.Picture
		dk.PictureSource = item.PictureSource

		for _, n := range item.NewsItem {
			dk.RelateNews = append(dk.RelateNews, &News{
				Title:   n.Title,
				Snippet: n.Snippet,
				Url:     n.Url,
				Source:  n.Source,
			})
		}

		ret = append(ret, dk)
	}

	return ret, tr.response, nil
}

func (dk *DailyKeyword) Marshal() string {
	b, _ := json.Marshal(&dk)
	return string(b)
}
