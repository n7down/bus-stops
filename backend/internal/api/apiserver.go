package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/bus-stops/backend/internal/api/busroutesapi"
)

type APIServer struct {
	busRoutesAPI *busroutesapi.BusRoutesAPI
}

func NewAPIServer(busRoutesAPI *busroutesapi.BusRoutesAPI) *APIServer {
	return &APIServer{
		busRoutesAPI: busRoutesAPI,
	}
}

func (s *APIServer) InitV1Routes(router *gin.Engine) error {
	v1 := router.Group("/api/v1")

	busGroup := v1.Group("/bus")
	{
		routesGroup := busGroup.Group("/routes")
		{
			routesGroup.POST("", s.busRoutesAPI.GetRoute)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return nil
}

func (g *APIServer) Run(router *gin.Engine, port string) error {
	err := http.ListenAndServe(port, router)
	if err != nil {
		return err
	}
	return nil
}
