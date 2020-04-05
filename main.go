package main

import (
	"io"

	"github.com/dtamura/hello-gin/lib/log"
	"github.com/dtamura/hello-gin/lib/tracing"
	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var tracer opentracing.Tracer
var logger log.Factory

func main() {
	// loggerの初期化
	logger1, _ := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapLogger := logger1.With(zap.String("service", "hello-gin"))
	logger = log.NewFactory(zapLogger)

	// OpenTracingの初期化
	var closer io.Closer
	tracer, closer = tracing.Init("hello-gin", logger)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer) // Jaeger tracer のグローバル変数を初期化

	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/healthz", healthz)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
