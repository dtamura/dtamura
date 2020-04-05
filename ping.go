package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func ping(c *gin.Context) {
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(c.Request.Header),
	)
	span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	hostname, _ := os.Hostname()

	span.SetTag("hostname", hostname) // Tagに"hello-to"をセット

	c.JSON(200, gin.H{
		"timestamp": time.Now(),
		"message":   "pong",
		"hostname":  hostname,
	})
}
