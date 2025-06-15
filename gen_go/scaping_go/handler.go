package main

import (
	"context"
)

// ScapingServiceImpl implements the last service interface defined in the IDL.
type ScapingServiceImpl struct{}

// Echo implements the ScapingServiceImpl interface.
func (s *ScapingServiceImpl) Echo(ctx context.Context, message string) (resp string, err error) {
	// TODO: Your code here...
	return
}
