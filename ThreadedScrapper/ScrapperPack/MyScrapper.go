package ScrapperPack

import (
	"fmt"
	"github.com/gocolly/colly/v2"
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





func ScrapInter(c *colly.Collector, EntryPoint string) {
	URLQueue := BFSQueue{}
	URLQueue.Push(EntryPoint)

	for len(URLQueue.MyUrls) > 0{
		URL := URLQueue.Pop()
		c.Visit(URL)


		ScrapeElements(c)
		//URLQueue = append(URLQueue, ScrapeURLS())
		c.Visit(URL)		
		//nextUrl := URLQueue.Pop()
		//Searched := InDataBase(nexturl)
		//for Searched == True{
			//Searched := InDataBase(nexturl)
		//}
	
	}

}


func ScrapeElements(c *colly.Collector) {

	//var AmazonProduct product;

	c.OnHTML("html", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

}
