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
		expectedRoute0 string
		expectedRoute1 string
		expectedRoute2 string
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
		assert.Equal(t, test.expectedRoute0, actualRoutes.Route0, fmt.Sprintf("%v should be %v", actualRoutes.Route0, test.expectedRoute0))
		assert.Equal(t, test.expectedRoute1, actualRoutes.Route1, fmt.Sprintf("%v should be %v", actualRoutes.Route1, test.expectedRoute1))
		assert.Equal(t, test.expectedRoute2, actualRoutes.Route2, fmt.Sprintf("%v should be %v", actualRoutes.Route2, test.expectedRoute2))
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
		assert.Equal(t, test.expectedRoute0, actualRoutes.Route0, fmt.Sprintf("%v should be %v", actualRoutes.Route0, test.expectedRoute0))
		assert.Equal(t, test.expectedRoute1, actualRoutes.Route1, fmt.Sprintf("%v should be %v", actualRoutes.Route1, test.expectedRoute1))
		assert.Equal(t, test.expectedRoute2, actualRoutes.Route2, fmt.Sprintf("%v should be %v", actualRoutes.Route2, test.expectedRoute2))
	}
}
