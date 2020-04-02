package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
)

func ping(c *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(c, "ping")
	defer span.Finish()

	hostname, _ := os.Hostname()

	span.SetTag("hostname", hostname) // Tagに"hello-to"をセット

	c.JSON(200, gin.H{
		"timestamp": time.Now(),
		"message":   "pong",
		"hostname":  hostname,
	})
}
