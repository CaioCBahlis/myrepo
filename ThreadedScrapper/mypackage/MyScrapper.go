package mypackage

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"time"
)

type BFSQueue struct{
	MyUrls []string
}

func (self *BFSQueue) Push(element string){
	self.MyUrls = append(self.MyUrls, element)
}

func (self *BFSQueue) Pop()string{
	elem := self.MyUrls[0]
	self.MyUrls = self.MyUrls[1:]
	return elem
}

//womp womp compiler


func ScrapInter(c *colly.Collector, EntryPoint string) {
	URLQueue := BFSQueue{}
	URLQueue.Push(EntryPoint)
	depth := 3

	//for len(URLQueue.MyUrls) > 0{
		//URL := URLQueue.Pop()
		//ScrapeElements(c)
		//c.Visit(URL)


		
		//URLQueue = append(URLQueue, ScrapeURLS())		
		//nextUrl := URLQueue.Pop()
		//Searched := InDataBase(nexturl)
		//for Searched == True{
			//Searched := InDataBase(nexturl)
		//}
	//}
	
	URL := URLQueue.Pop()
	

	for i := 0; i < depth; i++{
		
		ScrapeElements(c)

		c.OnHTML("html", func(divnext *colly.HTMLElement){
			URL = "https://www.amazon.com/" + divnext.ChildAttr("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator","href")
		  })

		c.Visit(URL)
	}

}

func ScrapeElements(c *colly.Collector){
	var MyItem product
	

	c.OnHTML("div[data-component-type=s-search-result]", func(e *colly.HTMLElement) {
	  
	  TitleParent := e.DOM.Find("div[data-cy=title-recipe]")
	  PriceParent := e.DOM.Find("span.a-price[data-a-size='xl']")
  
	  MyItem.Title = TitleParent.Find("span.a-size-medium.a-color-base.a-text-normal").Text()
	  MyItem.Price = PriceParent.Find("span.a-offscreen").Text()
	  MyItem.Reviews = e.ChildText("span.a-icon-alt")
	  
		
	  if MyItem.Title == ""{
		MyItem.Title =  TitleParent.Find("span.a-size-base-plus.a-color-base.a-text-normal").Text()
		MyItem.Price =  PriceParent.Find("span.a-offscreen").Text()
	  }
	  
	  PUrl, error  := e.DOM.Find("a.a-link-normal.s-no-outline").Attr("href")
	  if !error{
		fmt.Println("Error fetching Purl")
	  }
	  
	  MyItem.Purl = "https://www.amazon.com/" + PUrl
	  MyItem.Imgurl = e.ChildAttr("img.s-image","src")
	  MyItem.Seller = "Amazon"
	  
	
	  MyItem.Seller = "Amazon"
	  MyItem.Lupdate = time.DateOnly
	  if !DB_Search_and_Update(MyItem){
		DBinsert(MyItem)
	  }
	})

	
  
  }
