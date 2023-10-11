package utils

import "math"

func RoundAt(rate float64, size int) float64 {
	shift := math.Pow10(size)
	return math.Round(rate*(shift)) / (shift)
}

func CeilAt(rate float64, size int) float64 {
	shift := math.Pow10(size)
	return math.Ceil(rate*shift) / shift
}

func GetDistanceInMeters(lat1, lon1, lat2, lon2 float64) float64 {
	const EarthRadius = 6371000 // 지구 반지름 (미터)

	// 위도 및 경도를 라디안으로 변환
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// 위도 및 경도 간의 차이 계산
	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	// Haversine 공식을 사용하여 거리 계산
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := EarthRadius * c

	return distance
}
