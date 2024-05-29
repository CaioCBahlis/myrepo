package main

import (
	"ThreadedScrapper/ThreadedScrapper/mypackage"
	"fmt"
	"net/http"
	"github.com/gocolly/colly/v2"
	"github.com/imroc/req/v3"
)

//------------------------------------------------------
// Has to be Able to scrap multiple pages at once
//           MultiThreading on different links ????????????
//           Controllable Number of Threads //FIXED

// Has to Get Data and send it to database 
//        Get Product Data and send it to PostGres //FIXED
// Has to be Able to interact with database
//        Checks to Be Performed  
//        If product not in database, add //FIXED
// 		  if product in database, check the price for changes
// 		  if price is different, update price
//        if price is same, skip 
//                 Check Implementation
//                       Get Title and Search in DB
//                        if NOT IN db, add every attribute //FIXED
//                        if in DB compare prices
//                        true, skip, false, update
// Responsabilities of a Scrapper
// Scrapper Shouldnt be able to create threads, as req increases exponentially //FIXED
// Scrapper is going to be autonomous  //FIXED
// Design Choice made based on Number of interactions with PostGres  //FIXED
//


func main(){

	EntryLink := "https://www.amazon.com/s?k=esgrima&crid=2PK43XV5WY2RA&sprefix=esgri%2Caps%2C280&ref=nb_sb_noss_2"
	

	//MaxThreads := 1
	fakeBrowser := req.DefaultClient().ImpersonateFirefox()

	c := colly.NewCollector(
		colly.IgnoreRobotsTxt(),
		colly.UserAgent(fakeBrowser.Headers.Get("user-agent")),
	)

	c.SetClient(&http.Client{
		Transport: fakeBrowser.Transport,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnResponse(func(re *colly.Response) {
		fmt.Println(re.StatusCode)
	})


	mypackage.ScrapInter(c, EntryLink)
	
}
