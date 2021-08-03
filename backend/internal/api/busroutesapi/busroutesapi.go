package busroutesapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/bus-stops/backend/internal/api/busroutesapi/requests"
	"github.com/n7down/bus-stops/backend/internal/api/busroutesapi/responses"
	"github.com/n7down/bus-stops/backend/internal/persistence"
)

const (
	FIVE_MINUTES = time.Minute * 5
)

type BusRoutesAPI struct {
	persistence persistence.Persistence
}

func NewBusRoutesAPI(p persistence.Persistence) *BusRoutesAPI {
	return &BusRoutesAPI{
		persistence: p,
	}
}

func (a *BusRoutesAPI) GetRoute(c *gin.Context) {
	_, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req requests.GetRoutesRequest
		res responses.GetRoutesResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	routes := a.persistence.GetRoutes(req.Stop, req.Time)

	res = responses.GetRoutesResponse{
		Route1: routes.Route1,
		Route2: routes.Route2,
		Route3: routes.Route3,

		NextRoute1: routes.NextRoute1,
		NextRoute2: routes.NextRoute2,
		NextRoute3: routes.NextRoute3,
	}

	c.JSON(http.StatusOK, res)
}
