package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/imroc/req/v3"
)


type MyItem struct{
  Source string
  Price float32
  Rating string
  Category string
  Url string
}

/////////////////////////////////////////////////////////////////
// TO DO: Implement Random Impersonation
// TO DO: Implement 4 Sized Scrapping #FIXED
// TO DO: Implement DataBase
// TO DO: Implement Parsing based on characteristics #FIXED
// TO DO: Implement Multiple Pages  #FIXED
// TO DO: Implement GoRoutines
// TO DO: Implement Search Mode
/////////////////////////////////////////////////////////////////




func main(){

  fakeChrome := req.DefaultClient().ImpersonateFirefox()
  MAX_DEPTH := 1
  c := colly.NewCollector(
  colly.UserAgent(fakeChrome.Headers.Get("user-agent")),
  )

  c.SetClient(&http.Client{
  Transport: fakeChrome.Transport,
  })

  URL := "https://www.amazon.com/s?k=garrafa&crid=9QANVE34TVD7&sprefix=garra%2Caps%2C268&ref=nb_sb_noss_2"

  for i := 0; i != MAX_DEPTH; i++{
    Scrape(c)
    c.OnHTML("div.a-section.a-text-center.s-pagination-container", func(divnext *colly.HTMLElement){
      URL = "https://www.amazon.com/" + divnext.ChildAttr("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator","href")
  
    })

    time.Sleep(time.Duration(rand.Int31n(500)) * time.Millisecond)
    c.Visit(URL)

  }

  
 
  c.OnRequest(func(r *colly.Request) {
      log.Println("Visiting", r.URL)
    })
    
  c.Visit(URL)
  
}


func Scrape(c *colly.Collector){
  c.OnHTML("div.puis-card-container.s-card-container.s-overflow-hidden.aok-relative.puis-include-content-margin.puis.puis-vul871yo2a6ad24liafm30m715.s-latency-cf-section.puis-card-border", func(e *colly.HTMLElement) {


    e.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(i int, SingleGrid *colly.HTMLElement) {

      //fmt.Println(SingleGrid.Text)
      TitleParent := SingleGrid.DOM.Find("div[data-cy=title-recipe]")
      fmt.Println("Title:",  TitleParent.Find("span.a-size-medium.a-color-base.a-text-normal").Text())
      PriceParent := SingleGrid.DOM.Find("span.a-price[data-a-size='xl']")
      fmt.Println("Price:", PriceParent.Find("span.a-offscreen").Text())
      fmt.Println("Reviews:", SingleGrid.ChildText("span.a-icon-alt"))

    })
    
    
      
    e.ForEach("a.a-link-normal.s-no-outline", func(i int, singleURL *colly.HTMLElement) {
        fmt.Println("ProductUrl: https://www.amazon.com", singleURL.Attr("href"))

        fmt.Println("URL:", singleURL.ChildAttr("img.s-image","src"))
    })
      

    fmt.Println("--------------------------") 

    e.ForEach("div.a-section.a-spacing-small.puis-padding-left-small.puis-padding-right-small", func(i int, quadgrid *colly.HTMLElement) {
      TitleParent := quadgrid.DOM.Find("div[data-cy=title-recipe]")
      fmt.Println("Title:",  TitleParent.Find("span.a-size-base-plus.a-color-base.a-text-normal").Text())
      PriceParent := quadgrid.DOM.Find("span.a-price[data-a-size='xl']")
      fmt.Println("Price:", PriceParent.Find("span.a-offscreen").Text())
      fmt.Println("Reviews:", quadgrid.ChildText("span.a-icon-alt"))
    })

  })
}



