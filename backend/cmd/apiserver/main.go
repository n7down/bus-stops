package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/n7down/bus-stops/backend/internal/api"
	"github.com/n7down/bus-stops/backend/internal/api/busroutesapi"
	"github.com/n7down/bus-stops/backend/internal/persistence/stopsmock"
)

const (
	port = ":8080"
)

func main() {
	r := gin.Default()

	persistence := stopsmock.NewStops()

	busRoutesAPI := busroutesapi.NewBusRoutesAPI(persistence)

	apiServer := api.NewAPIServer(busRoutesAPI)

	apiServer.InitV1Routes(r)

	err := apiServer.Run(r, port)
	if err != nil {
		log.Fatal(err)
	}
}
