package main

import (
	"context"
)

// ScrapingServiceImpl implements the last service interface defined in the IDL.
type ScrapingServiceImpl struct{}

// Echo implements the ScrapingServiceImpl interface.
func (s *ScrapingServiceImpl) Echo(ctx context.Context, message string) (resp string, err error) {
	// TODO: Your code here...
	return
}
