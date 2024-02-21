package utils

import (
	"math"
	"strconv"
)

func ConvStrToFloat(value string) float32 {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return float32(val)
}

// haversine calculates the distance between two points on the Earth's surface using the Haversine formula.
// The input parameters are the latitude and longitude of the two points in decimal degrees.
// The function returns the distance between the two points in meters.
func haversine(lat1, long1, lat2, long2 float64) (distance float64) {
	const earthRadiusMetres = 6371000

	// convert degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	long1Rad := long1 * math.Pi / 180
	long2Rad := long2 * math.Pi / 180

	// calculate differences between latitudes and longitudes
	dLat := lat2Rad - lat1Rad
	dLong := long2Rad - long1Rad

	// apply the haversine formula
	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Pow(math.Sin(dLong/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// convert to meters
	distance = earthRadiusMetres * c

	return
}

type Geofence struct {
	CenterLat  float64 `json:"center_lat"`
	CenterLong float64 `json:"center_long"`
	Radius     float64 `json:"radius"`
	Rule       string  `json:"rule"`
}

// IsLocationInGeofence
//
// Parameters:
// lat (float64): The latitude of the location.
// long (float64): The longitude of the location.
// geofence (Geofence): The geofence to check against.
//
// Returns:
// status (bool): True if the location is in the geofence according to the rule, false otherwise.
// distance (float64): The distance between the location and the center of the geofence in meters.
func IsLocationInGeofence(
	lat, long float64,
	geofence Geofence,
) (status bool, distance float64) {
	distance = haversine(lat, long, geofence.CenterLat, geofence.CenterLong)
	status = false

	if geofence.Rule == "out" {
		// outside geofence
		if distance > geofence.Radius {
			status = true
		}
	} else {
		// inside geofence
		if distance <= geofence.Radius {
			status = true
		}
	}

	return
}
