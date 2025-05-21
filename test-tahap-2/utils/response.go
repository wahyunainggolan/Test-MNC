package utils

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, status int, message string) {
    c.AbortWithStatusJSON(status, gin.H{
        "status":  "FAILED",
        "message": message,
    })
}

func RespondSuccess(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, gin.H{
        "status": "SUCCESS",
        "result": data,
    })
}