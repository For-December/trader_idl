package main

import (
	"context"
	scraping "github.com/For-December/trader_idl/gen_go/scraping_go/kitex_gen/scraping"
)

// ScrapingServiceImpl implements the last service interface defined in the IDL.
type ScrapingServiceImpl struct{}

// Echo implements the ScrapingServiceImpl interface.
func (s *ScrapingServiceImpl) Echo(ctx context.Context, message string) (resp string, err error) {
	// TODO: Your code here...
	return
}

// GetBtcPrice implements the ScrapingServiceImpl interface.
func (s *ScrapingServiceImpl) GetBtcPrice(ctx context.Context, startTime int64, endTime int64, interval string) (resp []*scraping.PriceData, err error) {
	// TODO: Your code here...
	return
}

// GetSocialMediaData implements the ScrapingServiceImpl interface.
func (s *ScrapingServiceImpl) GetSocialMediaData(ctx context.Context, keyword string, startTime int64, endTime int64, limit int32) (resp []*scraping.TextData, err error) {
	// TODO: Your code here...
	return
}
