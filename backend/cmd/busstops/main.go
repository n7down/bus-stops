package main

import "github.com/n7down/bus-stops/backend/internal/persistence/bus"

func main() {
	busStops := bus.NewStops()
	busStops.Print()
}
