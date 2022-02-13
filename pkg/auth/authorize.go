package auth

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func Authenticate(ctx context.Context) (context.Context, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		if peer.AuthInfo == nil {
			return context.WithValue(ctx, subjectContextKey{}, ""), nil
		}
		tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
		subject := tlsInfo.State.VerifiedChains[0][0].Subject.CommonName
		ctx = context.WithValue(ctx, subjectContextKey{}, subject)

		return ctx, nil
	}
	return ctx, status.New(codes.Unknown, "couldn't find peer info").Err()
}

func subject(ctx context.Context) string {
	return ctx.Value(subjectContextKey{}).(string)
}

type subjectContextKey struct{}
