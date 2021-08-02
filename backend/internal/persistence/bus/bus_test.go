package bus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stops *Stops
)

func init() {
	stops = NewStops()
}

func TestGetStops(t *testing.T) {

	tests := []struct {
		stop           int
		stopTime       string
		expectedRoute0 string
		expectedRoute1 string
		expectedRoute2 string
	}{
		{
			0,
			"12:00AM",
			"12:00AM",
			"12:02AM",
			"12:04AM",
		},
		{
			1,
			"12:02AM",
			"12:02AM",
			"12:04AM",
			"12:06AM",
		},
		{
			2,
			"12:04AM",
			"12:04AM",
			"12:06AM",
			"12:08AM",
		},
		{
			3,
			"12:06AM",
			"12:06AM",
			"12:08AM",
			"12:10AM",
		},
		{
			4,
			"12:08AM",
			"12:08AM",
			"12:10AM",
			"12:12AM",
		},
	}

	for _, test := range tests {
		actualRouteMap := stops.Get(test.stop, test.stopTime)
		assert.Equal(t, test.expectedRoute0, actualRouteMap[0], fmt.Sprintf("%v should be %v", actualRouteMap[0], test.expectedRoute0))
		assert.Equal(t, test.expectedRoute1, actualRouteMap[1], fmt.Sprintf("%v should be %v", actualRouteMap[1], test.expectedRoute1))
		assert.Equal(t, test.expectedRoute2, actualRouteMap[2], fmt.Sprintf("%v should be %v", actualRouteMap[2], test.expectedRoute2))
	}
}

func TestGetStopsByFindingTheNextStop(t *testing.T) {

	tests := []struct {
		stop           int
		stopTime       string
		expectedRoute0 string
		expectedRoute1 string
		expectedRoute2 string
	}{
		{
			0,
			"12:00AM",
			"12:00AM",
			"12:02AM",
			"12:04AM",
		},
		{
			0,
			"12:01AM",
			"12:15AM",
			"12:17AM",
			"12:19AM",
		},
		{
			0,
			"03:01PM",
			"03:15PM",
			"03:17PM",
			"03:19PM",
		},
		{
			1,
			"12:03AM",
			"12:17AM",
			"12:19AM",
			"12:21AM",
		},
	}

	for _, test := range tests {
		actualRouteMap := stops.Get(test.stop, test.stopTime)
		assert.Equal(t, test.expectedRoute0, actualRouteMap[0], fmt.Sprintf("%v should be %v", actualRouteMap[0], test.expectedRoute0))
		assert.Equal(t, test.expectedRoute1, actualRouteMap[1], fmt.Sprintf("%v should be %v", actualRouteMap[1], test.expectedRoute1))
		assert.Equal(t, test.expectedRoute2, actualRouteMap[2], fmt.Sprintf("%v should be %v", actualRouteMap[2], test.expectedRoute2))
	}
}
