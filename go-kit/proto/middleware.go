package proto

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
)

// LoggerEndpointMiddleware returns go-kit middlewarefor which logs errors, panics & success (in debug)
func LoggerEndpointMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				r := recover()
				if r != nil { // panic
					lg.logger.Error(fmt.Errorf("Panic: %v", r), request, "duration", time.Since(begin))
				} else if err != nil { // error
					lg.logger.Error(err, request, "duration", time.Since(begin))
				} else { // success
					lg.logger.Debugw("Success", "payload", request, "duration", time.Since(begin))
				}
			}(time.Now())

			return next(ctx, request)
		}
	}
}
