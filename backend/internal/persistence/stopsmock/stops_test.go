package stopsmock

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
		expectedRoute1 string
		expectedRoute2 string
		expectedRoute3 string
	}{
		{
			1,
			"12:00AM",
			"12:00AM",
			"12:02AM",
			"12:04AM",
		},
		{
			2,
			"12:02AM",
			"12:02AM",
			"12:04AM",
			"12:06AM",
		},
		{
			3,
			"12:04AM",
			"12:04AM",
			"12:06AM",
			"12:08AM",
		},
		{
			4,
			"12:06AM",
			"12:06AM",
			"12:08AM",
			"12:10AM",
		},
		{
			5,
			"12:08AM",
			"12:08AM",
			"12:10AM",
			"12:12AM",
		},
	}

	for _, test := range tests {
		actualRoutes := stops.GetRoutes(test.stop, test.stopTime)
		assert.Equal(t, test.expectedRoute1, actualRoutes.Route1, fmt.Sprintf("%v should be %v", actualRoutes.Route1, test.expectedRoute1))
		assert.Equal(t, test.expectedRoute2, actualRoutes.Route2, fmt.Sprintf("%v should be %v", actualRoutes.Route2, test.expectedRoute2))
		assert.Equal(t, test.expectedRoute3, actualRoutes.Route3, fmt.Sprintf("%v should be %v", actualRoutes.Route3, test.expectedRoute3))
	}
}

func TestGetStopsByFindingTheNextStop(t *testing.T) {

	tests := []struct {
		stop           int
		stopTime       string
		expectedRoute1 string
		expectedRoute2 string
		expectedRoute3 string
	}{
		{
			1,
			"12:00AM",
			"12:00AM",
			"12:02AM",
			"12:04AM",
		},
		{
			1,
			"12:01AM",
			"12:15AM",
			"12:17AM",
			"12:19AM",
		},
		{
			1,
			"03:01PM",
			"03:15PM",
			"03:17PM",
			"03:19PM",
		},
		{
			2,
			"12:03AM",
			"12:17AM",
			"12:19AM",
			"12:21AM",
		},
	}

	for _, test := range tests {
		actualRoutes := stops.GetRoutes(test.stop, test.stopTime)
		assert.Equal(t, test.expectedRoute1, actualRoutes.Route1, fmt.Sprintf("%v should be %v", actualRoutes.Route1, test.expectedRoute1))
		assert.Equal(t, test.expectedRoute2, actualRoutes.Route2, fmt.Sprintf("%v should be %v", actualRoutes.Route2, test.expectedRoute2))
		assert.Equal(t, test.expectedRoute3, actualRoutes.Route3, fmt.Sprintf("%v should be %v", actualRoutes.Route3, test.expectedRoute3))
	}
}

func TestGetNextRouteStopsByFindingTheNextStop(t *testing.T) {

	tests := []struct {
		stop           int
		stopTime       string
		expectedRoute1 string
		expectedRoute2 string
		expectedRoute3 string

		expectedNextRoute1 string
		expectedNextRoute2 string
		expectedNextRoute3 string
	}{
		{
			1,
			"12:00AM",
			"12:00AM",
			"12:02AM",
			"12:04AM",
			"12:15AM",
			"12:17AM",
			"12:19AM",
		},
		{
			1,
			"12:01AM",
			"12:15AM",
			"12:17AM",
			"12:19AM",
			"12:30AM",
			"12:32AM",
			"12:34AM",
		},
		{
			1,
			"03:01PM",
			"03:15PM",
			"03:17PM",
			"03:19PM",
			"03:30PM",
			"03:32PM",
			"03:34PM",
		},
		{
			2,
			"12:03AM",
			"12:17AM",
			"12:19AM",
			"12:21AM",
			"12:32AM",
			"12:34AM",
			"12:36AM",
		},
	}

	for _, test := range tests {
		actualRoutes := stops.GetRoutes(test.stop, test.stopTime)
		assert.Equal(t, test.expectedRoute1, actualRoutes.Route1, fmt.Sprintf("%v should be %v", actualRoutes.Route1, test.expectedRoute1))
		assert.Equal(t, test.expectedRoute2, actualRoutes.Route2, fmt.Sprintf("%v should be %v", actualRoutes.Route2, test.expectedRoute2))
		assert.Equal(t, test.expectedRoute3, actualRoutes.Route3, fmt.Sprintf("%v should be %v", actualRoutes.Route3, test.expectedRoute3))
		assert.Equal(t, test.expectedNextRoute1, actualRoutes.NextRoute1, fmt.Sprintf("%v should be %v", actualRoutes.NextRoute1, test.expectedNextRoute1))
		assert.Equal(t, test.expectedNextRoute2, actualRoutes.NextRoute2, fmt.Sprintf("%v should be %v", actualRoutes.NextRoute2, test.expectedNextRoute2))
		assert.Equal(t, test.expectedNextRoute3, actualRoutes.NextRoute3, fmt.Sprintf("%v should be %v", actualRoutes.NextRoute3, test.expectedNextRoute3))
	}
}
