package main

import (
    "log"
    "net/http"
    "fmt"
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

func main(){
  fakeChrome := req.DefaultClient().ImpersonateFirefox()

  c := colly.NewCollector(
  colly.UserAgent(fakeChrome.Headers.Get("user-agent")),
  )

  c.SetClient(&http.Client{
  Transport: fakeChrome.Transport,
  })


  c.OnHTML("div.puis-card-container.s-card-container.s-overflow-hidden.aok-relative.puis-include-content-margin.puis.puis-vul871yo2a6ad24liafm30m715.s-latency-cf-section.puis-card-border", func(e *colly.HTMLElement) {
      //fmt.Println(e.Text)


      e.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(i int, h *colly.HTMLElement) {
        fmt.Println(h.Text)
      })

      e.ForEach("div.a-section.aok-relative.s-image-fixed-height", func(i int, im *colly.HTMLElement) {
        fmt.Println(im.ChildAttr("img.s-image","src"))
      })

      fmt.Println("--------------------------")
      
      
    })
    
  c.OnRequest(func(r *colly.Request) {
      log.Println("Visiting", r.URL)
    })
    
  c.Visit("https://www.amazon.com/s?rh=n%3A10&fs=true&ref=lp_10_sar")

}

