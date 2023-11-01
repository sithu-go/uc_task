package utils

const EarthRadiusKm = 6371 // Earth's radius in kilometers

// CalculateBounds calculates the latitude and longitude bounds based on the provided latitude, longitude, and radius (in kilometers)
func CalculateBounds(latitude, longitude, radiusKm float64) (latMin, latMax, lonMin, lonMax float64) {
	degreesPerKm := 1.0 / EarthRadiusKm

	latMin = latitude - (degreesPerKm * radiusKm)
	latMax = latitude + (degreesPerKm * radiusKm)
	lonMin = longitude - (degreesPerKm * radiusKm)
	lonMax = longitude + (degreesPerKm * radiusKm)

	return latMin, latMax, lonMin, lonMax
}
