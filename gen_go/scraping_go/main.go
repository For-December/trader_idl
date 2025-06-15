package main

import (
	scraping "github.com/For-December/trader_idl/gen_go/scraping_go/kitex_gen/scraping/scrapingservice"
	"log"
)

func main() {
	svr := scraping.NewServer(new(ScrapingServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
