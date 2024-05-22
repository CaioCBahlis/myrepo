package main


import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/imroc/req/v3"
) 


fakeChrome := req.DefaultClient().ImpersonateChrome()

c := colly.NewCollector(
 colly.MaxDepth(maxScrapeDepth),
 colly.UserAgent(fakeChrome.Headers.Get("user-agent")),
)
c.SetClient(&http.Client{
 Transport: fakeChrome.Transport,
})