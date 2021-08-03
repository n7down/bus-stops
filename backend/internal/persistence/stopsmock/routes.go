package stopsmock

import (
	"fmt"
	"time"

	"github.com/n7down/bus-stops/backend/internal/persistence"
)

type RoutesTimes struct {
	Route1 time.Time
	Route2 time.Time
	Route3 time.Time
}

func NewRoutesTimes(offset int) *RoutesTimes {

	t := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	// 12:00 am + offset
	timeOffset := time.Minute * time.Duration(offset)
	route1 := t.Add(timeOffset)

	// 12:02 am + offset
	timeOffset = time.Minute * time.Duration(offset+2)
	route2 := t.Add(timeOffset)

	// 12:04 am + offset
	timeOffset = time.Minute * time.Duration(offset+4)
	route3 := t.Add(timeOffset)

	return &RoutesTimes{
		Route1: route1,
		Route2: route2,
		Route3: route3,
	}
}

// func (r *Routes) Get(routeTime string) map[int]string {
func (r *RoutesTimes) Get(routeTime string) *persistence.Routes {
	twelveHourLayout := "03:04PM"

	newRouteTime, _ := time.Parse(twelveHourLayout, routeTime)

	routeSchedule1 := r.Route1
	routeSchedule2 := r.Route2
	routeSchedule3 := r.Route3

	diff := routeSchedule1.Sub(newRouteTime)

	for diff < 0 || diff > time.Hour*24 {

		// add 15 mins to each schedule
		routeSchedule1 = routeSchedule1.Add(time.Minute * 15)
		routeSchedule2 = routeSchedule2.Add(time.Minute * 15)
		routeSchedule3 = routeSchedule3.Add(time.Minute * 15)

		diff = routeSchedule1.Sub(newRouteTime)
	}

	route1 := routeSchedule1.Format(twelveHourLayout)
	route2 := routeSchedule2.Format(twelveHourLayout)
	route3 := routeSchedule3.Format(twelveHourLayout)

	nextRoute1 := routeSchedule1.Add(time.Minute * 15).Format(twelveHourLayout)
	nextRoute2 := routeSchedule2.Add(time.Minute * 15).Format(twelveHourLayout)
	nextRoute3 := routeSchedule3.Add(time.Minute * 15).Format(twelveHourLayout)

	ret := &persistence.Routes{
		Route1: route1,
		Route2: route2,
		Route3: route3,

		NextRoute1: nextRoute1,
		NextRoute2: nextRoute2,
		NextRoute3: nextRoute3,
	}

	return ret
}

func (r *RoutesTimes) Print() {
	fmt.Println(fmt.Sprintf("[1] %v", r.Route1))
	fmt.Println(fmt.Sprintf("[2] %v", r.Route2))
	fmt.Println(fmt.Sprintf("[3] %v", r.Route3))
}
