package handler

import (
	"SANS/internal/service"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Router  *gin.Engine
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
		Router:  gin.Default(),
	}
}

func (h *Handler) EndPoint() {
	h.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	//Daftar Fungsi
	// v1 := h.Router.Group("/v1")


	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	h.Router.Run(fmt.Sprintf(":%s", port))
}
