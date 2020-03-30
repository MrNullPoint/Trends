## Google Trends Unofficial Api

It's an unofficial api to get trending keyword from [google trends](https://trends.google.com/trends)

### Version

0.0.1

### Install

```shell
$ go get https://github.com/MrNullPoint/Trends
```

### Example Usage

- **Daily Keywords**: get daily keywords from google trends which you can set geo code such as `iso-3166-1` 2 bit code , default geo code is `US`

```go
func main() {
	tr := NewTrend()

	// you can set geo code here default is US
	tr.Geo("US")

	// you can set your own http client such as with a proxy client, like below
	// var transport http.Transport
	// var urli url.URL
	//
	// proxy, _ := urli.Parse("http://127.0.0.1:1080")
	// transport.Proxy = http.ProxyURL(proxy)
	//
	// client := new(http.Client)
	// client.Transport = &transport
	//
	// tr.Client(client)
	
	// get trends keywords from google
	keywords, resp, err := tr.DailyKeywordSearch()

	if err != nil {
		log.Fatal(err)
	}

	// you can use resp body to do something you like
	fmt.Println(resp.StatusCode)

	// you can use Trends keyword struct
	for _, k := range keywords {
		// you can get json string with Marshal()
		// fmt.Println(k.Marshal())

		// or you can get detail from keyword struct
		fmt.Println(k.Keyword) // keyword
		fmt.Println(k.Traffic) // search traffic
		fmt.Println(k.RelatedKeyword) // related keyword list
		fmt.Println(k.Picture) // picture url
		fmt.Println(k.PictureSource) // picture source
		fmt.Println(k.Link) // google trend link
		fmt.Println(k.Date) // keyword date

		for _, n := range k.RelateNews {
			fmt.Println(n.Title) // news title
			fmt.Println(n.Source) // news site name
			fmt.Println(n.Url) // news url
			fmt.Println(n.Snippet) // news snippet such as description
		}
	}
}
```

- **Rising Keywords**
- **Top Keywords**

### TODO

- [ ] RISING KEYWORD
- [ ] TOP KEYWORD
- [x] DAILY KEYWORD