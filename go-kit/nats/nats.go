package nats

import (
	"context"

	csxnats "github.com/csxbattle/nats"
	"github.com/zoobr/csxlib/logger"

	"github.com/go-kit/kit/endpoint"
	"github.com/nats-io/nats.go"
)

func SubscribeEndpoint(nc *csxnats.Conn, subject string, ctx context.Context, ep endpoint.Endpoint) {
	_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		_, err := ep(ctx, msg.Data)
		if err != nil {
			logger.Error(err, "queue", msg.Sub.Queue, "subject", msg.Subject)
		}
	})
	if err != nil {
		logger.Error(err, "url", nc.ConnectedUrl())
	}
}
