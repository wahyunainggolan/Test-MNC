package http

import (
    "github.com/gin-gonic/gin"
    "wallet-api/internal/service"
)

type Handler struct {
    UserService interface{} // simplify for demo
}

func NewHandler(r *gin.Engine, userService interface{}) {
    h := &Handler{UserService: userService}
    h.RegisterDashboardRoutes(r)
}