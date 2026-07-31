package config

import (
    "os"

    "github.com/gin-gonic/gin"
)

func GinMode() string {
    mode := os.Getenv("GIN_MODE")

    switch mode {
    case gin.ReleaseMode:
        return gin.ReleaseMode
    case gin.TestMode:
        return gin.TestMode
    default:
        return gin.DebugMode
    }
}