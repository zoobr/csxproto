package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/zoobr/csxlib/logger"
)

// LoggerEndpointMiddleware returns go-kit middlewarefor which logs errors, panics & success (in debug)
func LoggerEndpointMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				r := recover()
				if r != nil { // panic
					logger.Error(fmt.Errorf("Panic: %v", r), request, "duration", time.Since(begin))
				} else if err != nil { // error
					logger.Error(err, request, "duration", time.Since(begin))
				} else { // success
					logger.Debugw("Success", "payload", request, "duration", time.Since(begin))
				}
			}(time.Now())

			return next(ctx, request)
		}
	}
}

// ReqInfo logs with request ID getting from context
func ReqInfo(ctx context.Context, args ...interface{}) {
	reqID, err := getReqID(ctx)
	if err == nil {
		logger.Infof("Request ID: %s Payload: %s", reqID, args)
	}
}

// LoggerPathThrough returns a go-kit request function
// to add the request ID into context for path through logging
func LoggerPathThrough() httptransport.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		id, err := uuid.NewRandom()
		if err != nil {
			return ctx
		}
		reqID := id.String()
		return context.WithValue(ctx, "reqID", reqID)
	}
}

// getReqID returns request ID or error from context
// for path through logging
func getReqID(ctx context.Context) (reqID string, err error) {
	val := ctx.Value("reqID")
	if val == nil {
		return "", errors.New("can't get reqID from context")
	}
	return val.(string), nil
}
