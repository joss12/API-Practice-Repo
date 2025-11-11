package utils

import (
	"fmt"
	"math"
)

// RoundToDecimal rounds float to specified decimal places
func RoundToDecimal(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}

// Abs returns absolute value of float64
func Abs(value float64) float64 {
	return math.Abs(value)
}

// Clamp constrains value between min and max
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// NormalizeAngle normalizes angle to 0-360 degrees
func NormalizeAngle(degrees float64) float64 {
	for degrees < 0 {
		degrees += 360
	}
	for degrees >= 360 {
		degrees -= 360
	}
	return degrees
}

// FormatFloat formats float with specified decimal places
func FormatFloat(value float64, decimals int) string {
	return fmt.Sprintf("%."+fmt.Sprintf("%d", decimals)+"f", value)
}

// RadiansToDegrees converts radians to degrees
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// DegreesToRadians converts degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// SineDegrees calculates sine using degrees instead of radians
func SineDegrees(degrees float64) float64 {
	return math.Sin(DegreesToRadians(degrees))
}

// CosineDegrees calculates cosine using degrees instead of radians
func CosineDegrees(degrees float64) float64 {
	return math.Cos(DegreesToRadians(degrees))
}
