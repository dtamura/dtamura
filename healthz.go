package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func healthz(c *gin.Context) {
	hostname, _ := os.Hostname()
	c.JSON(200, gin.H{
		"timestamp": time.Now(),
		"status":    "OK",
		"message":   "i'm healthy",
		"hostname":  hostname,
	})
}
