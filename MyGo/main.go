package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
  "encoding/json"
  

	"github.com/gocolly/colly/v2"
	"github.com/imroc/req/v3"
)


type Item struct {
	Title      string `json:"title"`
	Price      string `json:"price"`
	Reviews    string `json:"reviews"`
	ProductUrl string `json:"product_url"`
	IMGUrl     string `json:"img_url"`
  seller     string `json:"seller"`
}


/////////////////////////////////////////////////////////////////
// TO DO: Implement Random Impersonation #Not Needed
// TO DO: Implement 4 Sized Scrapping #FIXED
// TO DO: Implement DataBase # Low Priority
// TO DO: Implement Parsing based on characteristics #FIXED
// TO DO: Implement Multiple Pages  #FIXED
// TO DO: Implement GoRoutines #To Be Implemented on Main File
// TO DO: Implement Search Mode #Fixed
// TO DO: Return type
/////////////////////////////////////////////////////////////////




func (){

  fmt.Println("Search on Amazon da Shoppe:")
  var searchreq string
  fmt.Scanln(&searchreq)
  URL := constructAmazonSearchURL(searchreq)

  fmt.Println("Max_Depth:")
  var MAX_DEPTH int
  fmt.Scan(&MAX_DEPTH)
  

  fakeChrome := req.DefaultClient().ImpersonateFirefox()
  
  c := colly.NewCollector(
  colly.UserAgent(fakeChrome.Headers.Get("user-agent")),
  )

  c.SetClient(&http.Client{
  Transport: fakeChrome.Transport,
  })

  for i := 0; i != MAX_DEPTH; i++{
    Scrape(c)

    c.OnHTML("div.a-section.a-text-center.s-pagination-container", func(divnext *colly.HTMLElement){
      URL = "https://www.amazon.com/" + divnext.ChildAttr("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator","href")
    })

    time.Sleep(time.Duration(rand.Int31n(500)) * time.Millisecond)
    c.Visit(URL)

  }

  c.OnResponse(func(r *colly.Response){
    fmt.Println(r.StatusCode)
  })
 
  c.OnRequest(func(r *colly.Request) {
      fmt.Println("Visiting", r.URL)
    })
    
  c.Visit(URL)
  
}


func Scrape(c *colly.Collector){
  
  c.OnHTML("div[data-component-type=s-search-result]", func(e *colly.HTMLElement) {
    var MyItem Item
    
    e.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(i int, SingleGrid *colly.HTMLElement) {

      //fmt.Println(SingleGrid.Text)
      TitleParent := SingleGrid.DOM.Find("div[data-cy=title-recipe]")
      PriceParent := SingleGrid.DOM.Find("span.a-price[data-a-size='xl']")

      MyItem.Title = TitleParent.Find("span.a-size-medium.a-color-base.a-text-normal").Text()
      MyItem.Price = PriceParent.Find("span.a-offscreen").Text()
      MyItem.Reviews = SingleGrid.ChildText("span.a-icon-alt")
    

    })
    
  
    e.ForEach("div.a-section.a-spacing-small.puis-padding-left-small.puis-padding-right-small", func(i int, quadgrid *colly.HTMLElement) {
      TitleParent := quadgrid.DOM.Find("div[data-cy=title-recipe]")
      PriceParent := quadgrid.DOM.Find("span.a-price[data-a-size='xl']")

      if MyItem.Title == ""{
        MyItem.Title =  TitleParent.Find("span.a-size-base-plus.a-color-base.a-text-normal").Text()
        MyItem.Price =  PriceParent.Find("span.a-offscreen").Text()
        MyItem.Reviews =  quadgrid.ChildText("span.a-icon-alt")
       }
      
    })

    e.ForEach("a.a-link-normal.s-no-outline", func(i int, singleURL *colly.HTMLElement) {
      MyItem.ProductUrl = "https://www.amazon.com" + singleURL.Attr("href")
      MyItem.IMGUrl = singleURL.ChildAttr("img.s-image","src")
      MyItem.seller = "Amazon"
      
  })

  JsonVal, err := json.Marshal(MyItem)
  if err != nil{
    fmt.Println("error converting object to json")
  }

  fmt.Println("--------------------------------------")
  fmt.Println(string(JsonVal))
  fmt.Println("--------------------------------------")

  })

}

func constructAmazonSearchURL(product string) string {
	baseURL := "https://www.amazon.com/s"
	params := url.Values{}
	params.Add("k", product)
	params.Add("ref", "nb_sb_noss_2")
	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}


