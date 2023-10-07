package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

func Logger() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			method := req.Spec().Procedure
			slog.Info("request received", slog.String("method", method))
			resp, err := next(ctx, req)
			if err != nil {
				slog.Error("request failed", slog.String("method", method), slog.Any("error", err))
			} else {
				slog.Info("request succeeded", slog.String("method", method))
			}
			return resp, err
		}
	}
}
