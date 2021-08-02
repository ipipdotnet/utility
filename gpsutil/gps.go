package gpsutil

type GPS struct {
	Longitude float64
	Latitude float64
}

func (g GPS) Distance(g2 GPS) float64 {
	return Distance(g, g2)
}

func Center(gs []GPS) GPS {
	return GPS{}
}