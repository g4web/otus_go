package internalgrpc

import (
	"context"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

var ErrPeerFromContext = status.Error(codes.Internal, "getting peer fail")

func RequestStatisticInterceptor(logger logger.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (response interface{}, err error) {
		p, ok := peer.FromContext(ctx)
		if !ok {
			logger.Error(ErrPeerFromContext.Error())

			return response, ErrPeerFromContext
		}

		start := time.Now()
		response, err = handler(ctx, req)

		logger.Info(" " + start.GoString() +
			" " + info.FullMethod +
			" " + p.Addr.String() +
			" " + time.Since(start).String())

		return response, err
	}
}
