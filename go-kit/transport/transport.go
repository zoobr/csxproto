package transport

import (
	"context"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/zoobr/csxlib/logger"
)

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

// ServerErrorLogger return HTTP server error logger for go-kit server
func ServerErrorLogger() httptransport.ServerOption {
	return httptransport.ServerErrorLogger(logger.GetLogger())
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
