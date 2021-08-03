package main

import "github.com/n7down/bus-stops/backend/internal/persistence/stopsmock"

func main() {
	busStops := stopsmock.NewStops()
	busStops.Print()
}
