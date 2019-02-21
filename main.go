package main

import (
	"context"

	"github.com/juju/zaputil/zapctx"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := zapctx.WithLevel(context.Background(), zapcore.DebugLevel)
	zapctx.Debug(ctx, "one")
	zapctx.Warn(ctx, "two")
}
