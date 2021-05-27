package util

import (
	"math"
)

// Equatorial radius
const a = 6378137.0

// Polar radius
const b = 6356752.314245

// central meridian of zone
const long0 = 121. / 180 * math.Pi

// scale along long0
const k0 = 0.9999

// delta x in meter
const dx = 250000.

//WGS84ToTWD97 convert wgs84 lat lon to twd 97 x,y
func WGS84ToTWD97(lat, lon float64) (xx, yy float64) {
	e := math.Pow((1 - math.Pow(b, 2)/math.Pow(a, 2)), 0.5)
	e2 := math.Pow(e, 2) / (1 - math.Pow(e, 2))
	n := (a - b) / (a + b)
	nu := a / math.Pow(1-math.Pow(e, 2)*(math.Sin(lat)*math.Sin(lat)), 0.5)
	p := lon - long0
	//log.Println(e, e2, n, nu, p)
	A := a * (1 - n + (5/4.0)*(math.Pow(n, 2)-math.Pow(n, 3)) + (81/64.0)*(math.Pow(n, 4)-math.Pow(n, 5)))
	B := (3 * a * n / 2.0) * (1 - n + (7/8.0)*(math.Pow(n, 2)-math.Pow(n, 3)) + (55/64.0)*(math.Pow(n, 4)-math.Pow(n, 5)))
	C := (15 * a * (math.Pow(n, 2)) / 16.0) * (1 - n + (3/4.0)*(math.Pow(n, 2)-math.Pow(n, 3)))
	D := (35 * a * (math.Pow(n, 3)) / 48.0) * (1 - n + (11/16.0)*(math.Pow(n, 2)-math.Pow(n, 3)))
	E := (315 * a * (math.Pow(n, 4)) / 51.0) * (1 - n)

	S := A*lat - B*math.Sin(2*lat) + C*math.Sin(4*lat) - D*math.Sin(6*lat) + E*math.Sin(8*lat)

	K1 := S * k0
	K2 := k0 * nu * math.Sin(2*lat) / 4.0
	K3 := (k0 * nu * math.Sin(lat) * math.Pow(math.Cos(lat), 3) / 24.0) * (5 - math.Pow(math.Tan(lat), 2) + 9*e2*(math.Pow(math.Cos(lat), 2)) + 4*(math.Pow(e2, 2))*(math.Pow(math.Cos(lat), 4)))

	yy = K1 + K2*(math.Pow(p, 2)) + K3*(math.Pow(p, 4))

	K4 := k0 * nu * math.Cos(lat)
	K5 := (k0 * nu * math.Pow(math.Cos(lat), 3) / 6.0) * (1 - math.Pow(math.Tan(lat), 2) + e2*math.Pow(math.Cos(lat), 2))

	xx = K4*p + K5*math.Pow(p, 3) + dx
	return xx, yy
}
