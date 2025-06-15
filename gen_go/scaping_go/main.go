package main

import (
	scaping "github.com/For-December/trader_idl/gen_go/scaping_go/kitex_gen/scaping/scapingservice"
	"log"
)

func main() {
	svr := scaping.NewServer(new(ScapingServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
