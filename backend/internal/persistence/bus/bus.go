package bus

import (
	"fmt"
	"time"
)

type Routes struct {
	RouteSchedule map[int]time.Time
}

func NewRoutes(offset int) *Routes {
	routeSchedule := make(map[int]time.Time)

	t := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	// 12:00 am + offset
	timeOffset := time.Minute * time.Duration(offset)
	routeSchedule[0] = t.Add(timeOffset)

	// 12:02 am + offset
	timeOffset = time.Minute * time.Duration(offset+2)
	routeSchedule[1] = t.Add(timeOffset)

	// 12:04 am + offset
	timeOffset = time.Minute * time.Duration(offset+4)
	routeSchedule[2] = t.Add(timeOffset)

	return &Routes{
		RouteSchedule: routeSchedule,
	}
}

func (r *Routes) Get(routeTime string) map[int]string {
	routeMap := make(map[int]string)
	twelveHourLayout := "03:04PM"

	newRouteTime, _ := time.Parse(twelveHourLayout, routeTime)

	routeSchedule0 := r.RouteSchedule[0]
	routeSchedule1 := r.RouteSchedule[1]
	routeSchedule2 := r.RouteSchedule[2]

	diff := routeSchedule0.Sub(newRouteTime)

	// FIXME: may need a limit on this so it doesnt go forever
	for diff < 0 {
		// add 15 mins to each schedule
		routeSchedule0 = routeSchedule0.Add(time.Minute * 15)
		routeSchedule1 = routeSchedule1.Add(time.Minute * 15)
		routeSchedule2 = routeSchedule2.Add(time.Minute * 15)

		diff = routeSchedule0.Sub(newRouteTime)
	}

	routeMap[0] = routeSchedule0.Format(twelveHourLayout)
	routeMap[1] = routeSchedule1.Format(twelveHourLayout)
	routeMap[2] = routeSchedule2.Format(twelveHourLayout)

	return routeMap
}

func (r *Routes) Print() {
	fmt.Println(fmt.Sprintf("[0] %v", r.RouteSchedule[0]))
	fmt.Println(fmt.Sprintf("[1] %v", r.RouteSchedule[1]))
	fmt.Println(fmt.Sprintf("[2] %v", r.RouteSchedule[2]))
}

type Stops struct {
	StopSchedule map[int]*Routes
}

func NewStops() *Stops {
	stopSchedule := make(map[int]*Routes)

	stopSchedule[0] = NewRoutes(0)
	stopSchedule[1] = NewRoutes(2)
	stopSchedule[2] = NewRoutes(4)
	stopSchedule[3] = NewRoutes(6)
	stopSchedule[4] = NewRoutes(8)
	stopSchedule[5] = NewRoutes(10)
	stopSchedule[6] = NewRoutes(12)
	stopSchedule[7] = NewRoutes(14)
	stopSchedule[8] = NewRoutes(16)
	stopSchedule[9] = NewRoutes(18)

	return &Stops{
		StopSchedule: stopSchedule,
	}
}

// TODO: make regex to check for time

func (s *Stops) Get(stop int, routeTime string) map[int]string {
	r := s.StopSchedule[stop].Get(routeTime)
	return r
}

func (s *Stops) Print() {
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("[%d]", i))
		s.StopSchedule[i].Print()
	}
}
