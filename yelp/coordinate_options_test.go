package yelp

import (
	"testing"

	"github.com/guregu/null"
)

// TestCoordinateOptions will check using location options.
func TestCoordinateOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		CoordinateOptions: &CoordinateOptions{
			null.FloatFrom(37.9),
			null.FloatFrom(-122.5),
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}
