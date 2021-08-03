package stopsmock

import (
	"fmt"
	"time"

	"github.com/n7down/bus-stops/backend/internal/persistence"
)

type RoutesTimes struct {
	// RouteSchedule map[int]time.Time
	Route0 time.Time
	Route1 time.Time
	Route2 time.Time
}

func NewRoutesTimes(offset int) *RoutesTimes {

	t := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	// 12:00 am + offset
	timeOffset := time.Minute * time.Duration(offset)
	route0 := t.Add(timeOffset)

	// 12:02 am + offset
	timeOffset = time.Minute * time.Duration(offset+2)
	route1 := t.Add(timeOffset)

	// 12:04 am + offset
	timeOffset = time.Minute * time.Duration(offset+4)
	route2 := t.Add(timeOffset)

	return &RoutesTimes{
		Route0: route0,
		Route1: route1,
		Route2: route2,
	}
}

// func (r *Routes) Get(routeTime string) map[int]string {
func (r *RoutesTimes) Get(routeTime string) *persistence.Routes {
	twelveHourLayout := "03:04PM"

	newRouteTime, _ := time.Parse(twelveHourLayout, routeTime)

	routeSchedule0 := r.Route0
	routeSchedule1 := r.Route1
	routeSchedule2 := r.Route2

	diff := routeSchedule0.Sub(newRouteTime)

	for diff < 0 || diff > time.Hour*24 {

		// add 15 mins to each schedule
		routeSchedule0 = routeSchedule0.Add(time.Minute * 15)
		routeSchedule1 = routeSchedule1.Add(time.Minute * 15)
		routeSchedule2 = routeSchedule2.Add(time.Minute * 15)

		diff = routeSchedule0.Sub(newRouteTime)
	}

	route0 := routeSchedule0.Format(twelveHourLayout)
	route1 := routeSchedule1.Format(twelveHourLayout)
	route2 := routeSchedule2.Format(twelveHourLayout)

	ret := &persistence.Routes{
		Route0: route0,
		Route1: route1,
		Route2: route2,
	}

	return ret
}

func (r *RoutesTimes) Print() {
	fmt.Println(fmt.Sprintf("[0] %v", r.Route0))
	fmt.Println(fmt.Sprintf("[1] %v", r.Route1))
	fmt.Println(fmt.Sprintf("[2] %v", r.Route2))
}
