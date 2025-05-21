package http

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func (h *Handler) RegisterDashboardRoutes(r *gin.Engine) {
    r.GET("/dashboard/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "Dashboard running",
        })
    })
}