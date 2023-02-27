package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	models "example/scrapper/utils/models"

	"github.com/gocolly/colly"
)

func GetCryptodata() [] models.Coin {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.MaxDepth(5),
	)
	var coinData []models.Coin
	//var selector = "#__next > div.sc-d89a36a8-0.hoSqxr.cmc-app-wrapper.cmc-app-wrapper--env-prod.cmc-theme--day > div.container.cmc-main-section > div > div.sc-31393a00-0.fljqdJ.cmc-view-all-coins > div > div.sc-a5cefde9-0.cUZGdI.cmc-table--sort-by__rank.cmc-table > div:nth-child(3) > div > table > tbody > tr"
	 var selector = ".cmc-table__table-wrapper-outer > div > table > tbody"
	
	 c.OnHTML(selector, func(e *colly.HTMLElement) {
		fmt.Println("Read")
		// log.Printf(e.ChildText(".currency-name-container"))
		// writer.Write([]string{
		// 	e.ChildText("td:nth-child(2)"),
		// 	e.ChildText("td:nth-child(3)"),
		// 	e.ChildAttr("a.price", "data-usd"),
		// 	e.ChildAttr("a.volume", "data-usd"),
		// 	e.ChildAttr(".market-cap", "data-usd"),
		// 	e.ChildText(".percent-1h"),
		// 	e.ChildText(".percent-24h"),
		// 	e.ChildText(".percent-7d"),
		// })
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
		
			data := models.Coin{
				Name:el.ChildText("td:nth-child(2)"),
				Symbol: el.ChildText("td:nth-child(3)"), 
				MarketCapUSD: el.ChildText("td:nth-child(4)"),
				PriceUSD: el.ChildText("td:nth-child(5)"),
			}
        	// fmt.Println(el.ChildText("td:nth-child(2)"))
			coinData = append(coinData, data)
    	})
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	return coinData
	
}

var baseSearchURL = "https://factba.se/json/json-transcript.php?q=&f=&dt=&p="
var baseTranscriptURL = "https://factba.se/transcript/"

type result struct {
	Slug string `json:"slug"`
	Date string `json:"date"`
}

type results struct {
	Data []*result `json:"data"`
}

type transcript struct {
	Speaker string
	Text    string
}


func SpeechText() {
	c := colly.NewCollector(
		colly.AllowedDomains("factba.se"),
	)

	d := c.Clone()

	d.OnHTML("body", func(e *colly.HTMLElement) {
		t := make([]transcript, 0)
		e.ForEach(".topic-media-row", func(_ int, el *colly.HTMLElement) {
			t = append(t, transcript{
				Speaker: el.ChildText(".speaker-label"),
				Text:    el.ChildText(".transcript-text-block"),
			})
		})
		jsonData, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return
		}
		ioutil.WriteFile(colly.SanitizeFileName(e.Request.Ctx.Get("date")+"_"+e.Request.Ctx.Get("slug"))+".json", jsonData, 0644)
	})

	stop := false
	c.OnResponse(func(r *colly.Response) {
		rs := &results{}
		err := json.Unmarshal(r.Body, rs)
		if err != nil || len(rs.Data) == 0 {
			stop = true
			return
		}
		for _, res := range rs.Data {
			u := baseTranscriptURL + res.Slug
			ctx := colly.NewContext()
			ctx.Put("date", res.Date)
			ctx.Put("slug", res.Slug)
			d.Request("GET", u, nil, ctx, nil)
		}
	})

	for i := 1; i < 10; i++ {
		if stop {
			break
		}
		if err := c.Visit(baseSearchURL + strconv.Itoa(i)); err != nil {
			fmt.Println("Error:", err)
			break
		}
	}

}