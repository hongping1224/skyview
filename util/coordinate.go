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
func TWD97ToWGS84(xx, yy float64) (lat, lon float64) {
	dy := float64(0)
	e := float64(math.Pow((1 - math.Pow(b, 2)/math.Pow(a, 2)), 0.5))
	xx -= dx
	yy -= dy
	// Calculate the Meridional Arc
	M := float64(yy / k0)

	// Calculate Footprint Latitude
	mu := M / (a * (1.0 - math.Pow(e, 2)/4.0 - 3*math.Pow(e, 4)/64.0 - 5*math.Pow(e, 6)/256.0))
	e1 := float64(1.0-math.Pow((1.0-math.Pow(e, 2)), 0.5)) / (1.0 + math.Pow((1.0-math.Pow(e, 2)), 0.5))
	J1 := float64(3*e1/2 - 27*math.Pow(e1, 3)/32.0)
	J2 := float64(21*math.Pow(e1, 2)/16 - 55*math.Pow(e1, 4)/32.0)
	J3 := float64(151 * math.Pow(e1, 3) / 96.0)
	J4 := float64(1097 * math.Pow(e1, 4) / 512.0)
	fp := float64(mu + J1*math.Sin(2*mu) + J2*math.Sin(4*mu) + J3*math.Sin(6*mu) + J4*math.Sin(8*mu))
	// Calculate Latitude and Longitude

	e2 := math.Pow((e * a / b), 2)
	C1 := math.Pow(e2*math.Cos(fp), 2)
	T1 := math.Pow(math.Tan(fp), 2)
	R1 := float64(a * (1 - math.Pow(e, 2)) / math.Pow((1-math.Pow(e, 2)*math.Pow(math.Sin(fp), 2)), (3.0/2.0)))
	N1 := float64(a / math.Pow((1-math.Pow(e, 2)*math.Pow(math.Sin(fp), 2)), 0.5))

	D := float64(xx / (N1 * k0))
	// 計算緯度
	Q1 := float64(N1 * math.Tan(fp) / R1)
	Q2 := float64(math.Pow(D, 2) / 2.0)
	Q3 := float64((5 + 3*T1 + 10*C1 - 4*math.Pow(C1, 2) - 9*e2) * math.Pow(D, 4) / 24.0)
	Q4 := float64((61 + 90*T1 + 298*C1 + 45*math.Pow(T1, 2) - 3*math.Pow(C1, 2) - 252*e2) * math.Pow(D, 6) / 720.0)
	lat = float64(fp - Q1*(Q2-Q3+Q4))

	// 計算經度
	Q5 := D
	Q6 := float64((1 + 2*T1 + C1) * math.Pow(D, 3) / 6.0)
	Q7 := float64((5 - 2*C1 + 28*T1 - 3*math.Pow(C1, 2) + 8*e2 + 24*math.Pow(T1, 2)) * math.Pow(D, 5) / 120.0)
	lon = float64(long0 + (Q5-Q6+Q7)/math.Cos(fp))

	lat = (lat * 180) / math.Pi //緯
	lon = (lon * 180) / math.Pi //經
	return
}
