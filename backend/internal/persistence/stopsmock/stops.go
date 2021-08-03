package stopsmock

import (
	"fmt"

	"github.com/n7down/bus-stops/backend/internal/persistence"
)

type Stops struct {
	StopSchedule map[int]*RoutesTimes
}

func NewStops() *Stops {
	stopSchedule := make(map[int]*RoutesTimes)

	stopSchedule[1] = NewRoutesTimes(0)
	stopSchedule[2] = NewRoutesTimes(2)
	stopSchedule[3] = NewRoutesTimes(4)
	stopSchedule[4] = NewRoutesTimes(6)
	stopSchedule[5] = NewRoutesTimes(8)
	stopSchedule[6] = NewRoutesTimes(10)
	stopSchedule[7] = NewRoutesTimes(12)
	stopSchedule[8] = NewRoutesTimes(14)
	stopSchedule[9] = NewRoutesTimes(16)
	stopSchedule[10] = NewRoutesTimes(18)

	return &Stops{
		StopSchedule: stopSchedule,
	}
}

func (s *Stops) GetRoutes(stop int, routeTime string) *persistence.Routes {
	r := s.StopSchedule[stop].Get(routeTime)
	return r
}

func (s *Stops) Print() {
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("[%d]", i))
		s.StopSchedule[i].Print()
	}
}
