package yelp

import (
	"errors"
	"fmt"

	"github.com/guregu/null"
)

// CoordinateOptions are used with complex searches for locations.
// The geographic coordinate format is defined as:
// ll=latitude,longitude,accuracy,altitude,altitude_accuracy
type CoordinateOptions struct {
	Latitude  null.Float // Latitude of geo-point to search near (required)
	Longitude null.Float // Longitude of geo-point to search near (required)
}

// getParameters will reflect over the values of the given
// struct, and provide a type appropriate set of querystring parameters
// that match the defined values.
func (o CoordinateOptions) getParameters() (params map[string]string, err error) {
	// coordinate requires at least a latitude and longitude - others are option
	if !o.Latitude.Valid || !o.Longitude.Valid {
		return nil, errors.New("latitude and longitude are required fields for a coordinate based search")
	}

	params = make(map[string]string)
	params["latitude"] = fmt.Sprintf("%v", o.Latitude.Float64)
	params["longitude"] = fmt.Sprintf("%v", o.Longitude.Float64)

	return params, nil
}
