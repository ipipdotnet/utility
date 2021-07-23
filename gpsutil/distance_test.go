package gpsutil

import "testing"

func TestDistance(t *testing.T) {

	t.Log(Distance(GPS{
		Longitude: 116.111,
		Latitude: 27.111,
	}, GPS{
		Longitude: 116.222,
		Latitude: 27.222,
	}))

}
